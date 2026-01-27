// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

// 导入OpenZeppelin的ERC721基础实现
import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/utils/Counters.sol";


/**
 * @title SimpleNFT
 * @dev 仅包含铸造和转移功能的基础NFT合约
 */
contract SimpleNFT is ERC721 {
    using Counters for Counters.Counter;
    
    // 使用计数器生成唯一的tokenId
    Counters.Counter private _tokenIdCounter;
    
    // 可选的铸造者地址（如果设定了只有特定地址可铸造）
    address public minter;
    
    // 铸造事件
    event Minted(address indexed to, uint256 indexed tokenId);
    
    /**
     * @dev 构造函数
     * @param name NFT集合名称
     * @param symbol NFT集合符号
     */
    constructor(string memory name, string memory symbol) ERC721(name, symbol) {
        // 默认铸造者为合约部署者
        minter = msg.sender;
    }
    
    /**
     * @dev 修改铸造者权限（仅当前minter可调用）
     * @param newMinter 新的铸造者地址
     */
    function setMinter(address newMinter) external {
        require(msg.sender == minter, "Only minter can change minter");
        minter = newMinter;
    }
    
    /**
     * @dev 铸造NFT（仅铸造者可调用）
     * @param to NFT接收者地址
     * @return 新铸造的NFT的tokenId
     */
    function mint(address to) external returns (uint256) {
        // 检查调用者是否有铸造权限
        require(msg.sender == minter, "Only minter can mint");
        
        // 生成新的tokenId
        uint256 tokenId = _tokenIdCounter.current();
        _tokenIdCounter.increment();
        
        // 铸造NFT
        _safeMint(to, tokenId);
        
        // 触发铸造事件
        emit Minted(to, tokenId);
        
        return tokenId;
    }
    
    /**
     * @dev 批量铸造NFT
     * @param to NFT接收者地址
     * @param amount 铸造数量
     */
    function mintBatch(address to, uint256 amount) external {
        require(msg.sender == minter, "Only minter can mint");
        
        for (uint256 i = 0; i < amount; i++) {
            uint256 tokenId = _tokenIdCounter.current();
            _tokenIdCounter.increment();
            _safeMint(to, tokenId);
            emit Minted(to, tokenId);
        }
    }
    
    /**
     * @dev 获取已铸造的NFT总数
     * @return 当前NFT总量
     */
    function totalSupply() external view returns (uint256) {
        return _tokenIdCounter.current();
    }
    
    // 注意：transferFrom和safeTransferFrom函数已从ERC721父合约继承
    // 用户可以直接使用这些标准ERC721转移函数
}