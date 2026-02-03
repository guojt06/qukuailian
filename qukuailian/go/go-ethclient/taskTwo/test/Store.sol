pragma solidity ^0.8.26;

contract Store {
    event ItemSet(bytes32 key, bytes32 value);

    string public version;
    mapping(bytes32 => bytes32) public items;

    constructor(string memory _version) {
        version = _version;
    }

    function setItem(bytes32 key, bytes32 value) external {
        items[key] = value;
        emit ItemSet(key, value);
    }
    // 添加一个新的函数来获取version
    // 注意：public变量会自动生成getter函数，这里只是为了演示如何添加自定义函数
    function getVersion() external view returns (string memory) {
        return version;
    }

    // 或者可以添加一个获取version和当前区块号的功能
    function getVersionWithBlock() external view returns (string memory, uint256) {
        return (version, block.number);
    }
}