// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract RomanNumerals {
    // 罗马数字符号映射
    struct RomanSymbol {
        uint256 value;
        string symbol;
    }

    // 罗马数字符号表（从大到小排列）
    RomanSymbol[] private romanSymbols;

    constructor() {
        // 初始化符号表
        romanSymbols.push(RomanSymbol(1000, "M"));
        romanSymbols.push(RomanSymbol(900, "CM"));
        romanSymbols.push(RomanSymbol(500, "D"));
        romanSymbols.push(RomanSymbol(400, "CD"));
        romanSymbols.push(RomanSymbol(100, "C"));
        romanSymbols.push(RomanSymbol(90, "XC"));
        romanSymbols.push(RomanSymbol(50, "L"));
        romanSymbols.push(RomanSymbol(40, "XL"));
        romanSymbols.push(RomanSymbol(10, "X"));
        romanSymbols.push(RomanSymbol(9, "IX"));
        romanSymbols.push(RomanSymbol(5, "V"));
        romanSymbols.push(RomanSymbol(4, "IV"));
        romanSymbols.push(RomanSymbol(1, "I"));
    }


    /**
     * @dev 优化的整数转罗马数字实现（使用字符串拼接）
     * @param num 要转换的整数（1-3999）
     * @return 罗马数字字符串
     */
    function intToRomanOptimized(uint256 num) public pure returns (string memory) {
        require(num > 0 && num < 4000, "Number must be between 1 and 3999");

        // 使用字符串直接拼接（更简洁）
        string memory result;
        uint256 remainder = num;

        // 处理千位
        if (remainder >= 1000) {
            uint256 thousands = remainder / 1000;
            for (uint256 i = 0; i < thousands; i++) {
                result = string(abi.encodePacked(result, "M"));
            }
            remainder %= 1000;
        }

        // 处理百位
        if (remainder >= 100) {
            uint256 hundreds = remainder / 100;
            if (hundreds == 9) {
                result = string(abi.encodePacked(result, "CM"));
            } else if (hundreds >= 5) {
                result = string(abi.encodePacked(result, "D"));
                for (uint256 i = 0; i < hundreds - 5; i++) {
                    result = string(abi.encodePacked(result, "C"));
                }
            } else if (hundreds == 4) {
                result = string(abi.encodePacked(result, "CD"));
            } else {
                for (uint256 i = 0; i < hundreds; i++) {
                    result = string(abi.encodePacked(result, "C"));
                }
            }
            remainder %= 100;
        }

        //MMMCMXCVII

        // 处理十位
        if (remainder >= 10) {
            uint256 tens = remainder / 10;
            if (tens == 9) {
                result = string(abi.encodePacked(result, "XC"));
            } else if (tens >= 5) {
                result = string(abi.encodePacked(result, "L"));
                for (uint256 i = 0; i < tens - 5; i++) {
                    result = string(abi.encodePacked(result, "X"));
                }
            } else if (tens == 4) {
                result = string(abi.encodePacked(result, "XL"));
            } else {
                for (uint256 i = 0; i < tens; i++) {
                    result = string(abi.encodePacked(result, "X"));
                }
            }
            remainder %= 10;
        }

        // 处理个位
        if (remainder > 0) {
            if (remainder == 9) {
                result = string(abi.encodePacked(result, "IX"));
            } else if (remainder >= 5) {
                result = string(abi.encodePacked(result, "V"));
                for (uint256 i = 0; i < remainder - 5; i++) {
                    result = string(abi.encodePacked(result, "I"));
                }
            } else if (remainder == 4) {
                result = string(abi.encodePacked(result, "IV"));
            } else {
                for (uint256 i = 0; i < remainder; i++) {
                    result = string(abi.encodePacked(result, "I"));
                }
            }
        }

        return result;
    }
}