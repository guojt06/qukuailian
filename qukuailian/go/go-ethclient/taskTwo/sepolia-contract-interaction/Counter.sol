// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/**
 * @title 简单计数器合约
 * @dev 演示基本的合约交互
 */
contract Counter {
    uint256 private count;
    address public owner;

    // 事件
    event CountIncremented(address indexed caller, uint256 newValue);
    event CountReset(address indexed caller);
    event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);

    // 构造函数
    constructor() {
        count = 0;
        owner = msg.sender;
    }

    // 修饰器：仅所有者可调用
    modifier onlyOwner() {
        require(msg.sender == owner, "Only owner can call this function");
        _;
    }

    /**
     * @dev 获取当前计数
     * @return 当前计数值
     */
    function getCount() public view returns (uint256) {
        return count;
    }

    /**
     * @dev 增加计数
     * @return 新的计数值
     */
    function increment() public returns (uint256) {
        count += 1;
        emit CountIncremented(msg.sender, count);
        return count;
    }

    /**
     * @dev 减少计数
     * @return 新的计数值
     */
    function decrement() public returns (uint256) {
        require(count > 0, "Counter cannot be negative");
        count -= 1;
        return count;
    }

    /**
     * @dev 重置计数器（仅所有者）
     */
    function reset() public onlyOwner {
        count = 0;
        emit CountReset(msg.sender);
    }

    /**
     * @dev 转移所有权
     * @param newOwner 新的所有者地址
     */
    function transferOwnership(address newOwner) public onlyOwner {
        require(newOwner != address(0), "New owner cannot be zero address");
        emit OwnershipTransferred(owner, newOwner);
        owner = newOwner;
    }

    /**
     * @dev 设置特定值（仅所有者）
     * @param newCount 新的计数值
     */
    function setCount(uint256 newCount) public onlyOwner {
        count = newCount;
    }
}