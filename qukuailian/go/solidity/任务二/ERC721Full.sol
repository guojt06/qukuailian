// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

/**
 * @title 完整的 ERC721 实现
 * @dev 实现 ERC721 标准，包含元数据、枚举、安全转移等功能
 */
contract ERC721Full {
    // =============================================================
    //                          事件
    // =============================================================

    event Transfer(
        address indexed from,
        address indexed to,
        uint256 indexed tokenId
    );

    event Approval(
        address indexed owner,
        address indexed approved,
        uint256 indexed tokenId
    );

    event ApprovalForAll(
        address indexed owner,
        address indexed operator,
        bool approved
    );

    // =============================================================
    //                          存储
    // =============================================================

    // Token 名称
    string private _name;

    // Token 符号
    string private _symbol;

    // tokenId -> 所有者地址
    mapping(uint256 => address) private _owners;

    // 地址 -> 拥有的 Token 数量
    mapping(address => uint256) private _balances;

    // tokenId -> 被授权的地址（单个授权）
    mapping(uint256 => address) private _tokenApprovals;

    // 所有者地址 -> 操作员地址 -> 是否授权所有代币
    mapping(address => mapping(address => bool)) private _operatorApprovals;

    // tokenId -> 元数据 URI
    mapping(uint256 => string) private _tokenURIs;

    // 基础 URI
    string private _baseURI;

    // Token ID 计数器
    uint256 private _nextTokenId = 1;

    // =============================================================
    //                          构造函数
    // =============================================================

    /**
     * @dev 初始化合约，设置 Token 名称和符号
     * @param name_ Token 名称
     * @param symbol_ Token 符号
     */
    constructor(string memory name_, string memory symbol_) {
        _name = name_;
        _symbol = symbol_;
    }

    // =============================================================
    //                          ERC721 标准函数
    // =============================================================

    /**
     * @dev 返回 Token 名称
     */
    function name() public view returns (string memory) {
        return _name;
    }

    /**
     * @dev 返回 Token 符号
     */
    function symbol() public view returns (string memory) {
        return _symbol;
    }

    /**
     * @dev 返回指定地址拥有的 Token 数量
     * @param owner 要查询的地址
     */
    function balanceOf(address owner) public view returns (uint256) {
        require(owner != address(0), "ERC721: address zero is not a valid owner");
        return _balances[owner];
    }

    /**
     * @dev 返回指定 Token ID 的所有者
     * @param tokenId 要查询的 Token ID
     */
    function ownerOf(uint256 tokenId) public view returns (address) {
        address owner = _ownerOf(tokenId);
        require(owner != address(0), "ERC721: invalid token ID");
        return owner;
    }

    /**
     * @dev 批准另一个地址转移指定的 Token
     * @param to 被授权的地址
     * @param tokenId 要授权的 Token ID
     */
    function approve(address to, uint256 tokenId) public {
        address owner = ownerOf(tokenId);
        require(to != owner, "ERC721: approval to current owner");

        require(
            msg.sender == owner || isApprovedForAll(owner, msg.sender),
            "ERC721: approve caller is not token owner or approved for all"
        );

        _approve(to, tokenId);
    }

    /**
     * @dev 获取指定 Token 的被授权地址
     * @param tokenId 要查询的 Token ID
     */
    function getApproved(uint256 tokenId) public view returns (address) {
        _requireMinted(tokenId);
        return _tokenApprovals[tokenId];
    }

    /**
     * @dev 设置或取消设置操作员（管理所有代币）
     * @param operator 操作员地址
     * @param approved 是否授权
     */
    function setApprovalForAll(address operator, bool approved) public {
        _setApprovalForAll(msg.sender, operator, approved);
    }

    /**
     * @dev 查询操作员是否被授权管理所有代币
     * @param owner 所有者地址
     * @param operator 操作员地址
     */
    function isApprovedForAll(address owner, address operator) public view returns (bool) {
        return _operatorApprovals[owner][operator];
    }

    /**
     * @dev 转移 Token（调用者需是所有者或被授权者）
     * @param from 当前所有者
     * @param to 接收者
     * @param tokenId 要转移的 Token ID
     */
    function transferFrom(address from, address to, uint256 tokenId) public {
        require(_isApprovedOrOwner(msg.sender, tokenId), "ERC721: caller is not token owner or approved");
        _transfer(from, to, tokenId);
    }

    /**
     * @dev 安全转移 Token（会检查接收者是否能处理 ERC721 Token）
     * @param from 当前所有者
     * @param to 接收者
     * @param tokenId 要转移的 Token ID
     */
    function safeTransferFrom(address from, address to, uint256 tokenId) public {
        safeTransferFrom(from, to, tokenId, "");
    }

    /**
     * @dev 安全转移 Token（带数据）
     * @param from 当前所有者
     * @param to 接收者
     * @param tokenId 要转移的 Token ID
     * @param data 附加数据
     */
    function safeTransferFrom(address from, address to, uint256 tokenId, bytes memory data) public {
        require(_isApprovedOrOwner(msg.sender, tokenId), "ERC721: caller is not token owner or approved");
        _safeTransfer(from, to, tokenId, data);
    }

    // =============================================================
    //                          ERC721Metadata 扩展
    // =============================================================

    /**
     * @dev 返回指定 Token 的 URI
     * @param tokenId 要查询的 Token ID
     */
    function tokenURI(uint256 tokenId) public view returns (string memory) {
        _requireMinted(tokenId);

        string memory _tokenURI = _tokenURIs[tokenId];
        string memory base = _baseURI;

        // 如果设置了基础 URI，但没有设置 Token URI
        if (bytes(base).length > 0) {
            // 如果有基础 URI，但没有特定 Token URI，使用默认格式
            if (bytes(_tokenURI).length == 0) {
                return string(abi.encodePacked(base, _toString(tokenId)));
            }
            // 如果有基础 URI 和特定 Token URI，合并它们
            return string(abi.encodePacked(base, _tokenURI));
        }

        return _tokenURI;
    }

    /**
     * @dev 设置基础 URI
     * @param baseURI_ 基础 URI
     */
    function _setBaseURI(string memory baseURI_) internal {
        _baseURI = baseURI_;
    }

    /**
     * @dev 设置 Token URI
     * @param tokenId Token ID
     * @param _tokenURI Token URI
     */
    function _setTokenURI(uint256 tokenId, string memory _tokenURI) internal {
        _requireMinted(tokenId);
        _tokenURIs[tokenId] = _tokenURI;
    }

    // =============================================================
    //                          ERC721Enumerable 扩展
    // =============================================================

    /**
     * @dev 返回已铸造的 Token 总数
     */
    function totalSupply() public view returns (uint256) {
        return _nextTokenId - 1;
    }

    // =============================================================
    //                          铸造函数
    // =============================================================

    /**
     * @dev 铸造新的 Token
     * @param to 接收者地址
     * @param tokenURI_ Token URI（可选）
     * @return 新铸造的 Token ID
     */
    function mint(address to, string memory tokenURI_) public returns (uint256) {
        return _mint(to, tokenURI_);
    }

    /**
     * @dev 安全铸造新的 Token
     * @param to 接收者地址
     * @param tokenURI_ Token URI（可选）
     * @return 新铸造的 Token ID
     */
    function safeMint(address to, string memory tokenURI_) public returns (uint256) {
        return _safeMint(to, tokenURI_);
    }

    /**
     * @dev 批量铸造 Token
     * @param to 接收者地址数组
     * @param tokenURIs_ Token URI 数组
     */
    function batchMint(address[] memory to, string[] memory tokenURIs_) public {
        require(to.length == tokenURIs_.length, "ERC721: arrays length mismatch");

        for (uint256 i = 0; i < to.length; i++) {
            _mint(to[i], tokenURIs_[i]);
        }
    }

    // =============================================================
    //                          内部函数
    // =============================================================

    /**
     * @dev 内部函数：铸造 Token
     */
    function _mint(address to, string memory tokenURI_) internal returns (uint256) {
        require(to != address(0), "ERC721: mint to the zero address");

        uint256 tokenId = _nextTokenId;
        _nextTokenId++;

        _balances[to] += 1;
        _owners[tokenId] = to;

        if (bytes(tokenURI_).length > 0) {
            _setTokenURI(tokenId, tokenURI_);
        }

        emit Transfer(address(0), to, tokenId);

        return tokenId;
    }

    /**
     * @dev 内部函数：安全铸造 Token
     */
    function _safeMint(address to, string memory tokenURI_) internal returns (uint256) {
        uint256 tokenId = _mint(to, tokenURI_);
        _checkOnERC721Received(address(0), to, tokenId, "");
        return tokenId;
    }

    /**
     * @dev 内部函数：转移 Token
     */
    function _transfer(address from, address to, uint256 tokenId) internal {
        require(ownerOf(tokenId) == from, "ERC721: transfer from incorrect owner");
        require(to != address(0), "ERC721: transfer to the zero address");

        // 清除之前的授权
        _approve(address(0), tokenId);

        _balances[from] -= 1;
        _balances[to] += 1;
        _owners[tokenId] = to;

        emit Transfer(from, to, tokenId);
    }

    /**
     * @dev 内部函数：安全转移 Token
     */
    function _safeTransfer(address from, address to, uint256 tokenId, bytes memory data) internal {
        _transfer(from, to, tokenId);
        _checkOnERC721Received(from, to, tokenId, data);
    }

    /**
     * @dev 内部函数：授权 Token
     */
    function _approve(address to, uint256 tokenId) internal {
        _tokenApprovals[tokenId] = to;
        emit Approval(ownerOf(tokenId), to, tokenId);
    }

    /**
     * @dev 内部函数：设置操作员授权
     */
    function _setApprovalForAll(address owner, address operator, bool approved) internal {
        require(owner != operator, "ERC721: approve to caller");
        _operatorApprovals[owner][operator] = approved;
        emit ApprovalForAll(owner, operator, approved);
    }

    /**
     * @dev 内部函数：检查是否已铸造
     */
    function _requireMinted(uint256 tokenId) internal view {
        require(_exists(tokenId), "ERC721: invalid token ID");
    }

    /**
     * @dev 内部函数：检查 Token 是否存在
     */
    function _exists(uint256 tokenId) internal view returns (bool) {
        return _owners[tokenId] != address(0);
    }

    /**
     * @dev 内部函数：获取所有者（不进行零地址检查）
     */
    function _ownerOf(uint256 tokenId) internal view returns (address) {
        return _owners[tokenId];
    }

    /**
     * @dev 内部函数：检查调用者是否有权限
     */
    function _isApprovedOrOwner(address spender, uint256 tokenId) internal view returns (bool) {
        address owner = ownerOf(tokenId);
        return (spender == owner ||
        getApproved(tokenId) == spender ||
            isApprovedForAll(owner, spender));
    }

    /**
     * @dev 内部函数：检查接收者是否能处理 ERC721 Token
     */
    function _checkOnERC721Received(
        address from,
        address to,
        uint256 tokenId,
        bytes memory data
    ) private {
        if (to.code.length > 0) {
            try IERC721Receiver(to).onERC721Received(msg.sender, from, tokenId, data) returns (bytes4 retval) {
                require(retval == IERC721Receiver.onERC721Received.selector, "ERC721: transfer to non ERC721Receiver implementer");
            } catch (bytes memory reason) {
                if (reason.length == 0) {
                    revert("ERC721: transfer to non ERC721Receiver implementer");
                } else {
                    assembly {
                        revert(add(32, reason), mload(reason))
                    }
                }
            }
        }
    }

    /**
     * @dev 内部函数：将 uint256 转换为 string
     */
    function _toString(uint256 value) internal pure returns (string memory) {
        if (value == 0) {
            return "0";
        }
        uint256 temp = value;
        uint256 digits;
        while (temp != 0) {
            digits++;
            temp /= 10;
        }
        bytes memory buffer = new bytes(digits);
        while (value != 0) {
            digits -= 1;
            buffer[digits] = bytes1(uint8(48 + uint256(value % 10)));
            value /= 10;
        }
        return string(buffer);
    }
}

// =============================================================
//                          IERC721Receiver 接口
// =============================================================

interface IERC721Receiver {
    function onERC721Received(
        address operator,
        address from,
        uint256 tokenId,
        bytes calldata data
    ) external returns (bytes4);
}