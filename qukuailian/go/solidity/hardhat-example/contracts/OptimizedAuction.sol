// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721Receiver.sol";
import "@openzeppelin/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";

// import {console} from "hardhat/console.sol";

/**
 * @title NFT拍卖合约
 * @notice 实现NFT的拍卖功能，支持ETH和ERC20参与竞拍，价格以USD计价
 */
contract OptimizedAuction is
    IERC721Receiver,
    Initializable,
    UUPSUpgradeable,
    OwnableUpgradeable,
    ReentrancyGuardUpgradeable
{
    // ============ 常量定义 ============
    uint256 private constant PRICE_FEED_DECIMALS = 8;
    uint256 private constant USD_DECIMALS = 6;
    uint256 private constant ETH_DECIMALS = 18;
    
    // ============ 结构体定义 ============
    struct Auction {
        address seller;         // 卖家
        uint256 startTime;      // 开始时间
        uint256 duration;       // 拍卖时长
        uint256 startPrice;     // 起拍价格（代币单位）
        bool ended;             // 是否结束
        address highestBidder;  // 最高出价者
        uint256 highestBid;     // 最高出价（代币单位）
        address nftContract;    // NFT合约地址
        uint256 tokenId;        // NFT ID
        address payToken;       // 支付代币类型
        address factory;        // 工厂合约地址
        uint256 auctionId;      // 拍卖ID
    }
    
    // ============ 状态变量 ============
    Auction private auction;
    mapping(address => uint256) private tokenPrices; // 代币价格（8位小数）
    
    // ============ 事件定义 ============
    event AuctionCreated(
        uint256 indexed auctionId,
        address indexed seller,
        address indexed nftContract,
        uint256 tokenId,
        uint256 startPrice,
        address payToken,
        uint256 startTime,
        uint256 endTime
    );
    
    event NewBid(
        uint256 indexed auctionId,
        address indexed bidder,
        uint256 amount,
        uint256 usdValue
    );
    
    event AuctionEnded(
        uint256 indexed auctionId,
        address indexed winner,
        uint256 amount,
        uint256 usdValue
    );
    
    event BidRefunded(
        uint256 indexed auctionId,
        address indexed bidder,
        uint256 amount
    );
    
    event NFTTransferred(
        address indexed nftContract,
        uint256 indexed tokenId,
        address indexed to
    );

    // ============ 初始化函数 ============
    function initialize(
        address _usdcAddress,
        address _seller,
        address _nftContract,
        uint256 _duration,
        uint256 _startPrice,
        uint256 _tokenId,
        address _payToken,
        address _factory,
        uint256 _auctionId
    ) public initializer {
        __Ownable_init(_seller); // 卖家作为初始所有者
        __UUPSUpgradeable_init();
        __ReentrancyGuard_init();

        // 参数验证
        require(_seller != address(0), "Invalid seller");
        require(_nftContract != address(0), "Invalid NFT contract");
        require(_duration >= 300, "Duration too short"); // 至少5分钟
        require(_startPrice > 0, "Start price must be positive");
        require(_factory != address(0), "Invalid factory");
        
        // 验证NFT所有权和授权
        IERC721 nft = IERC721(_nftContract);
        require(nft.ownerOf(_tokenId) == _seller, "Seller not owner");
        require(
            nft.getApproved(_tokenId) == address(this) || 
            nft.isApprovedForAll(_seller, address(this)),
            "Contract not approved"
        );

        // 初始化价格
        _initTokenPrices(_usdcAddress);
        
        // 验证支付代币是否支持
        if (_payToken != address(0)) {
            require(tokenPrices[_payToken] > 0, "Token not supported");
        }

        // 计算结束时间
        uint256 endTime = block.timestamp + _duration;
        
        // 创建拍卖
        auction = Auction({
            seller: _seller,
            startTime: block.timestamp,
            duration: _duration,
            startPrice: _startPrice,
            ended: false,
            highestBidder: address(0),
            highestBid: 0,
            nftContract: _nftContract,
            tokenId: _tokenId,
            payToken: _payToken,
            factory: _factory,
            auctionId: _auctionId
        });

        emit AuctionCreated(
            _auctionId,
            _seller,
            _nftContract,
            _tokenId,
            _startPrice,
            _payToken,
            block.timestamp,
            endTime
        );
    }

    // ============ 外部函数 ============

    /**
     * @dev 竞拍出价
     * @param _amount 出价金额（必须使用拍卖指定的代币）
     */
    function placeBid(uint256 _amount) external payable nonReentrant {
        require(!auction.ended, "Auction ended");
        require(block.timestamp < auction.startTime + auction.duration, "Auction expired");
        require(_amount > 0, "Bid amount must be positive");
        require(msg.sender != auction.seller, "Seller cannot bid");
        
        address payToken = auction.payToken;
        
        // 验证支付
        if (payToken == address(0)) {
            require(msg.value == _amount, "ETH amount mismatch");
        } else {
            require(msg.value == 0, "ERC20 bid should not send ETH");
            require(
                IERC20(payToken).allowance(msg.sender, address(this)) >= _amount,
                "Insufficient allowance"
            );
            
            bool success = IERC20(payToken).transferFrom(msg.sender, address(this), _amount);
            require(success, "ERC20 transfer failed");
        }

        // 计算并验证USD价值
        uint256 currentHighestUSD = _getCurrentHighestUSD();
        uint256 bidUSD = _calculateUSDValue(payToken, _amount);
        require(bidUSD > currentHighestUSD, "Bid too low");

        // 退还上一个出价者
        if (auction.highestBidder != address(0)) {
            _refundBidder(auction.highestBidder, payToken, auction.highestBid);
        }

        // 更新拍卖状态
        auction.highestBidder = msg.sender;
        auction.highestBid = _amount;

        emit NewBid(auction.auctionId, msg.sender, _amount, bidUSD);
    }

    /**
     * @dev 结束拍卖
     */
    function endAuction() external nonReentrant {
        require(!auction.ended, "Auction already ended");
        require(block.timestamp >= auction.startTime + auction.duration, "Auction not finished");
        
        auction.ended = true;
        IERC721 nft = IERC721(auction.nftContract);

        if (auction.highestBidder != address(0)) {
            // 转移NFT给中标者
            nft.safeTransferFrom(address(this), auction.highestBidder, auction.tokenId);
            emit NFTTransferred(auction.nftContract, auction.tokenId, auction.highestBidder);

            // 计算手续费
            (uint256 platformFee, uint256 sellerAmount) = _calculateFees(auction.highestBid);
            
            // 支付给卖家
            _transferFunds(auction.seller, auction.payToken, sellerAmount);
            
            // 支付平台手续费
            address platformAddress = _getPlatformAddress();
            if (platformFee > 0) {
                _transferFunds(platformAddress, auction.payToken, platformFee);
            }

            emit AuctionEnded(
                auction.auctionId,
                auction.highestBidder,
                auction.highestBid,
                _calculateUSDValue(auction.payToken, auction.highestBid)
            );
        } else {
            // 无人出价，退回NFT
            nft.safeTransferFrom(address(this), auction.seller, auction.tokenId);
            emit NFTTransferred(auction.nftContract, auction.tokenId, auction.seller);
            
            emit AuctionEnded(auction.auctionId, address(0), 0, 0);
        }
    }

    // ============ 视图函数 ============

    /**
     * @dev 获取拍卖信息
     */
    function getAuctionInfo() external view returns (
        address seller,
        uint256 startTime,
        uint256 endTime,
        uint256 startPrice,
        bool ended,
        address highestBidder,
        uint256 highestBid,
        address nftContract,
        uint256 tokenId,
        address payToken,
        uint256 auctionId
    ) {
        return (
            auction.seller,
            auction.startTime,
            auction.startTime + auction.duration,
            auction.startPrice,
            auction.ended,
            auction.highestBidder,
            auction.highestBid,
            auction.nftContract,
            auction.tokenId,
            auction.payToken,
            auction.auctionId
        );
    }

    /**
     * @dev 获取当前最高出价的USD价值
     */
    function getCurrentHighestUSD() external view returns (uint256) {
        return _getCurrentHighestUSD();
    }

    /**
     * @dev 获取代币价格
     */
    function getTokenPrice(address _token) external view returns (uint256) {
        return tokenPrices[_token];
    }

    /**
     * @dev 获取最小出价金额（USD）
     */
    function getMinBidUSD() external view returns (uint256) {
        uint256 currentUSD = _getCurrentHighestUSD();
        return currentUSD + (currentUSD * 5) / 100; // 加5%
    }

    /**
     * @dev 获取拍卖剩余时间
     */
    function getTimeRemaining() external view returns (uint256) {
        uint256 endTime = auction.startTime + auction.duration;
        if (auction.ended || block.timestamp >= endTime) {
            return 0;
        }
        return endTime - block.timestamp;
    }

    // ============ 内部函数 ============

    /**
     * @dev 初始化代币价格
     */
    function _initTokenPrices(address _usdcAddress) private {
        // ETH/USD: 假设 $2500
        tokenPrices[address(0)] = 2500 * (10 ** PRICE_FEED_DECIMALS);
        
        // USDC/USD: $1
        tokenPrices[_usdcAddress] = 1 * (10 ** PRICE_FEED_DECIMALS);
    }

    /**
     * @dev 计算USD价值
     */
    function _calculateUSDValue(address _token, uint256 _amount) private view returns (uint256) {
        uint256 price = tokenPrices[_token];
        require(price > 0, "Token price not found");
        
        // 获取代币精度
        uint256 tokenDecimals = _token == address(0) ? 
            ETH_DECIMALS : 
            IERC20Metadata(_token).decimals();
        
        // 计算公式: amount * price / 10^(tokenDecimals + priceDecimals - usdDecimals)
        uint256 denominator = 10 ** (tokenDecimals + PRICE_FEED_DECIMALS - USD_DECIMALS);
        
        return (_amount * price) / denominator;
    }

    /**
     * @dev 获取当前最高USD价值
     */
    function _getCurrentHighestUSD() private view returns (uint256) {
        if (auction.highestBidder != address(0)) {
            return _calculateUSDValue(auction.payToken, auction.highestBid);
        }
        return _calculateUSDValue(auction.payToken, auction.startPrice);
    }

    /**
     * @dev 退还出价者资金
     */
    function _refundBidder(address _bidder, address _token, uint256 _amount) private {
        _transferFunds(_bidder, _token, _amount);
        emit BidRefunded(auction.auctionId, _bidder, _amount);
    }

    /**
     * @dev 资金转账
     */
    function _transferFunds(address _to, address _token, uint256 _amount) private {
        require(_to != address(0), "Invalid recipient");
        require(_amount > 0, "Amount must be positive");
        
        if (_token == address(0)) {
            // 使用call代替transfer
            (bool success, ) = payable(_to).call{value: _amount}("");
            require(success, "ETH transfer failed");
        } else {
            bool success = IERC20(_token).transfer(_to, _amount);
            require(success, "ERC20 transfer failed");
        }
    }

    /**
     * @dev 计算手续费
     */
    function _calculateFees(uint256 _totalAmount) private returns (uint256 platformFee, uint256 sellerAmount) {
        (bool success, bytes memory data) = auction.factory.call(
            abi.encodeWithSignature("calculateFee(uint256)", _totalAmount)
        );
        require(success, "Fee calculation failed");
        
        platformFee = abi.decode(data, (uint256));
        require(platformFee <= _totalAmount, "Invalid fee amount");
        
        sellerAmount = _totalAmount - platformFee;
    }

    /**
     * @dev 获取平台地址
     */
    function _getPlatformAddress() private returns (address) {
        (bool success, bytes memory data) = auction.factory.call(
            abi.encodeWithSignature("getPlatformFeeRecipient()")
        );
        require(success, "Failed to get platform address");
        return abi.decode(data, (address));
    }

    // ============ ERC721接收函数 ============
    
    function onERC721Received(
        address,
        address,
        uint256,
        bytes calldata
    ) external pure override returns (bytes4) {
        return this.onERC721Received.selector;
    }

    // ============ UUPS升级函数 ============
    
    function _authorizeUpgrade(address newImplementation) internal override onlyOwner {}

    // ============ 接收ETH函数 ============
    
    receive() external payable {}
}