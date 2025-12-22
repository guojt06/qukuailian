// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/**
 * @title 标准 ERC20 代币合约
 * @dev 实现 ERC20 标准，包含增发功能
 */
contract MyToken {
    // 代币基本信息
    string public name;
    string public symbol;
    uint8 public decimals;
    uint256 public totalSupply;

    // 合约所有者
    address public owner;

    // 账户余额映射
    mapping(address => uint256) private _balances;

    // 授权额度映射
    mapping(address => mapping(address => uint256)) private _allowances;

    // 事件定义
    event Transfer(address indexed from, address indexed to, uint256 value);
    event Approval(address indexed owner, address indexed spender, uint256 value);
    event Mint(address indexed to, uint256 amount);
    event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);

    /**
     * @dev 构造函数，初始化代币
     * @param _name 代币名称
     * @param _symbol 代币符号
     * @param _decimals 小数位数
     * @param _initialSupply 初始供应量
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
        _mint(msg.sender, _initialSupply);
    }

    // 修饰器：只有合约所有者可以调用
    modifier onlyOwner() {
        require(msg.sender == owner, "MyToken: caller is not the owner");
        _;
    }

    /**
     * @dev 查询账户余额
     * @param account 要查询的账户地址
     * @return 账户余额
     */
    function balanceOf(address account) public view returns (uint256) {
        return _balances[account];
    }

    /**
     * @dev 转账函数
     * @param recipient 接收者地址
     * @param amount 转账金额
     * @return 是否成功
     */
    function transfer(address recipient, uint256 amount) public returns (bool) {
        _transfer(msg.sender, recipient, amount);
        return true;
    }

    /**
     * @dev 查询授权额度
     * @param owner 代币所有者
     * @param spender 被授权者
     * @return 授权额度
     */
    function allowance(address owner, address spender) public view returns (uint256) {
        return _allowances[owner][spender];
    }

    /**
     * @dev 授权函数
     * @param spender 被授权者地址
     * @param amount 授权金额
     * @return 是否成功
     */
    function approve(address spender, uint256 amount) public returns (bool) {
        _approve(msg.sender, spender, amount);
        return true;
    }

    /**
     * @dev 代扣转账函数（需要预先授权）
     * @param sender 发送者地址
     * @param recipient 接收者地址
     * @param amount 转账金额
     * @return 是否成功
     */
    function transferFrom(
        address sender,
        address recipient,
        uint256 amount
    ) public returns (bool) {
        _transfer(sender, recipient, amount);

        // 检查授权额度并更新
        uint256 currentAllowance = _allowances[sender][msg.sender];
        require(currentAllowance >= amount, "MyToken: transfer amount exceeds allowance");
        _approve(sender, msg.sender, currentAllowance - amount);

        return true;
    }

    /**
     * @dev 增发代币（仅所有者可调用）
     * @param to 接收增发代币的地址
     * @param amount 增发数量
     */
    function mint(address to, uint256 amount) public onlyOwner {
        _mint(to, amount);
    }

    /**
     * @dev 转移合约所有权
     * @param newOwner 新的所有者地址
     */
    function transferOwnership(address newOwner) public onlyOwner {
        require(newOwner != address(0), "MyToken: new owner is the zero address");
        emit OwnershipTransferred(owner, newOwner);
        owner = newOwner;
    }

    /**
     * @dev 内部转账实现
     */
    function _transfer(
        address sender,
        address recipient,
        uint256 amount
    ) internal {
        require(sender != address(0), "MyToken: transfer from the zero address");
        require(recipient != address(0), "MyToken: transfer to the zero address");

        uint256 senderBalance = _balances[sender];
        require(senderBalance >= amount, "MyToken: transfer amount exceeds balance");

        _balances[sender] = senderBalance - amount;
        _balances[recipient] += amount;

        emit Transfer(sender, recipient, amount);
    }

    /**
     * @dev 内部授权实现
     */
    function _approve(
        address owner_,
        address spender,
        uint256 amount
    ) internal {
        require(owner_ != address(0), "MyToken: approve from the zero address");
        require(spender != address(0), "MyToken: approve to the zero address");

        _allowances[owner_][spender] = amount;
        emit Approval(owner_, spender, amount);
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
     * @dev 销毁代币
     * @param amount 销毁数量
     */
    function burn(uint256 amount) public {
        _burn(msg.sender, amount);
    }

    /**
     * @dev 代扣销毁（需要预先授权）
     * @param account 销毁代币的账户
     * @param amount 销毁数量
     */
    function burnFrom(address account, uint256 amount) public {
        uint256 currentAllowance = allowance(account, msg.sender);
        require(currentAllowance >= amount, "MyToken: burn amount exceeds allowance");
        _approve(account, msg.sender, currentAllowance - amount);
        _burn(account, amount);
    }

    /**
     * @dev 内部销毁实现
     */
    function _burn(address account, uint256 amount) internal {
        require(account != address(0), "MyToken: burn from the zero address");

        uint256 accountBalance = _balances[account];
        require(accountBalance >= amount, "MyToken: burn amount exceeds balance");

        _balances[account] = accountBalance - amount;
        totalSupply -= amount;

        emit Transfer(account, address(0), amount);
    }

    /**
     * @dev 批量转账（提高效率）
     * @param recipients 接收者地址数组
     * @param amounts 转账金额数组
     */
    function batchTransfer(
        address[] memory recipients,
        uint256[] memory amounts
    ) public returns (bool) {
        require(recipients.length == amounts.length, "MyToken: arrays length mismatch");

        uint256 totalAmount = 0;
        for (uint256 i = 0; i < amounts.length; i++) {
            totalAmount += amounts[i];
        }

        require(_balances[msg.sender] >= totalAmount, "MyToken: insufficient balance");

        for (uint256 i = 0; i < recipients.length; i++) {
            _transfer(msg.sender, recipients[i], amounts[i]);
        }

        return true;
    }

    /**
     * @dev 增加授权额度（安全方式）
     */
    function increaseAllowance(address spender, uint256 addedValue) public returns (bool) {
        _approve(msg.sender, spender, _allowances[msg.sender][spender] + addedValue);
        return true;
    }

    /**
     * @dev 减少授权额度（安全方式）
     */
    function decreaseAllowance(address spender, uint256 subtractedValue) public returns (bool) {
        uint256 currentAllowance = _allowances[msg.sender][spender];
        require(currentAllowance >= subtractedValue, "MyToken: decreased allowance below zero");
        _approve(msg.sender, spender, currentAllowance - subtractedValue);
        return true;
    }
}