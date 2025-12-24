// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/**
 * @title 标准 ERC20 代币合约
 * @dev 实现 ERC20 标准，包含增发功能
 */
//erc20代币合约理解： 就是初始化币种后，有多个地址将币进行分配，其中合约拥有者还可以给某一个地址进行授权额度

//a：transferFrom接口如何使用
//  这是转账接口 a->b  其中ERC20接口是将币的数据进行变更 a-数量   b+数量
//  ERC721接口是 a->b 是将token ID的所有者地址进行变更

//b：ERC721中的授权接口跟ERC20有何不同  approve
//  erc20接口
//

//c：safeTransfer等带safe前缀的接口提供了什么安全措施
contract MyToken {
    // 定义代币基本信息
    string public name; //代币名称
    string public symbol; //代币符号
    uint8 public decimals; // 小数位数
    uint256 public totalSupply; // 代币总数

    address public owner;  //合约创建者
    // 存储映射
    mapping(address => uint256) private _balances;  //余额映射
    mapping(address => mapping(address => uint256)) private _allowances;  //授权地址-》映射

    // 事件定义
    event Transfer(address indexed from, address indexed to, uint256 value);
    event Approval(address indexed owner, address indexed spender, uint256 value);
    event Mint(address indexed to, uint256 amount);
    event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);

    /**
     * @dev 构造函数，初始化代币
     */
    constructor(
        string memory _name,
        string memory _symbol,
        uint8 _decimals,
        uint256 _initialSupply
    ) {
        name = _name;
        symbol = _symbol;
        decimals = _decimals;
        owner = msg.sender;

        // 将初始供应量分配给合约部署者
        _mint(msg.sender, _initialSupply * 10 ** _decimals);
    }

    // 修饰器：只有合约所有者可以调用
    modifier onlyOwner() {
        require(msg.sender == owner, "MyToken: caller is not the owner");
        _;
    }

    /**
     * @dev 查询账户余额
     */
    function balanceOf(address account) public view returns (uint256) {
        return _balances[account];
    }

    /**
     * @dev 向指定地址转账
     * @param to 接收地址
     * @param amount 转账金额
     */
    function transfer(address to, uint256 amount) public returns (bool) {
        require(to != address(0), "MyToken: transfer to the zero address");
        require(_balances[msg.sender] >= amount, "MyToken: insufficient balance");

        _balances[msg.sender] -= amount;
        _balances[to] += amount;

        emit Transfer(msg.sender, to, amount);
        return true;
    }

    /**
     * @dev 授权额度给指定地址
     * @param spender 被授权地址
     * @param amount 授权金额
     */
    function approve(address spender, uint256 amount) public returns (bool) {
        require(spender != address(0), "MyToken: approve to the zero address");

        _allowances[msg.sender][spender] = amount;
        emit Approval(msg.sender, spender, amount);
        return true;
    }

    /**
     * @dev 查询授权额度
     */
    function allowance(address tokenOwner, address spender) public view returns (uint256) {
        return _allowances[tokenOwner][spender];
    }

    /**
     * @dev 使用授权额度进行转账
     * @param from 扣款地址
     * @param to 收款地址
     * @param amount 转账金额
     */
    function transferFrom(address from, address to, uint256 amount) public returns (bool) {
        require(from != address(0), "MyToken: transfer from the zero address");
        require(to != address(0), "MyToken: transfer to the zero address");
        require(_balances[from] >= amount, "MyToken: insufficient balance");
        require(_allowances[from][msg.sender] >= amount, "MyToken: transfer amount exceeds allowance");

        _balances[from] -= amount;
        _balances[to] += amount;
        _allowances[from][msg.sender] -= amount;

        emit Transfer(from, to, amount);
        return true;
    }

    /**
     * @dev 允许合约所有者增发代币（公开函数）
     */
    function mint(address to, uint256 amount) public onlyOwner returns (bool) {
        _mint(to, amount);
        return true;
    }

    /**
     * @dev 内部增发实现
     */
    function _mint(address account, uint256 amount) internal {
        require(account != address(0), "MyToken: mint to the zero address");

        totalSupply += amount;
        _balances[account] += amount;

        emit Mint(account, amount);
        emit Transfer(address(0), account, amount);
    }

    /**
     * @dev 转移合约所有权
     */
    function transferOwnership(address newOwner) public onlyOwner {
        require(newOwner != address(0), "MyToken: new owner is the zero address");

        emit OwnershipTransferred(owner, newOwner);
        owner = newOwner;
    }
}