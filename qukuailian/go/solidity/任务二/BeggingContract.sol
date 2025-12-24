// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

contract BeggingContract {
    // 记录每个捐赠者的捐赠金额
    mapping(address => uint256) public donations;

    // 合约所有者地址
    address public owner;

    // 总捐赠金额
    uint256 public totalDonations;

    // 事件：记录捐赠信息
    event Donated(address indexed donor, uint256 amount);
    event Withdrawn(address indexed owner, uint256 amount);

    // 构造函数：设置合约所有者为部署者
    constructor() {
        owner = msg.sender;
    }

    // 修饰器：限制只有所有者可以调用
    modifier onlyOwner() {
        require(msg.sender == owner, "Only owner can call this function");
        _;
    }

    // donate 函数：接收以太币捐赠
    function donate() external payable {
        require(msg.value > 0, "Donation amount must be greater than 0");

        // 更新捐赠记录
        donations[msg.sender] += msg.value;
        totalDonations += msg.value;

        // 触发捐赠事件
        emit Donated(msg.sender, msg.value);
    }

    // withdraw 函数：提取所有资金（仅限所有者）
    function withdraw() external onlyOwner {
        // 获取合约当前余额
        uint256 balance = address(this).balance;
        require(balance > 0, "No funds to withdraw");

        // 重置总捐赠金额（可选，根据需求决定）
        // totalDonations = 0;

        // 提取资金到所有者地址
        payable(owner).transfer(balance);

        // 触发提取事件
        emit Withdrawn(owner, balance);
    }

    // getDonation 函数：查询指定地址的捐赠金额
    function getDonation(address donor) external view returns (uint256) {
        return donations[donor];
    }

    // 获取合约当前余额
    function getContractBalance() external view returns (uint256) {
        return address(this).balance;
    }

    // 接收以太币的回退函数（可选）
    receive() external payable {
        // 直接调用 donate 函数处理
        donations[msg.sender] += msg.value;
        totalDonations += msg.value;
        emit Donated(msg.sender, msg.value);
    }
}