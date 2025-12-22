// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

//反转字符串 (Reverse String)
//题目描述：反转一个字符串。输入 "abcde"，输出 "edcba"

contract StringReverse {


    // 字符串反转
    function setString(string memory name) public pure returns ( string memory ) {
        // 将字符串转换为字节数组
        bytes memory strBytes = bytes(name);
        bytes memory reversed = new bytes(strBytes.length);

        // 反转字节数组
        for (uint256 i = 0; i < strBytes.length; i++) {
            reversed[i] = strBytes[strBytes.length - 1 - i];
        }

        // 将字节数组转回字符串
        return string(reversed);
    }

}