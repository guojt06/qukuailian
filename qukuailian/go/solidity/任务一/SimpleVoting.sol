// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

//✅ 创建一个名为Voting的合约，包含以下功能：
//一个mapping来存储候选人的得票数
//一个vote函数，允许用户投票给某个候选人
//一个getVotes函数，返回某个候选人的得票数
//一个resetVotes函数，重置所有候选人的得票数

contract Voting {
    //一个mapping来储存候选人的得票数
    mapping(string => uint256) public poll;

    //创建一个候选人list
    string[] public namelist = ["zhangsan","lisi","wangwu"];

    // vote函数：允许用户投票给某个候选人
    function vote(string memory name) public {
        require(bytes(name).length > 0, "dsfds");

        // 增加票数
        poll[name] += 1;

    }

    // getVotes函数：返回某个候选人的得票数
    function getVotes(string memory name) public view returns (uint256) {
        return poll[name];
    }

    // resetVotes函数：重置所有候选人的得票数
    function resetVotes() public {
        for (uint256 i = 0; i < namelist.length; i++) {
            poll[namelist[i]] = 0;
        }
    }
}