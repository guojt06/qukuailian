// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract RomanToInteger {
    // 罗马数字到整数的映射
    mapping(bytes1 => uint256) private romanValues;

    // 特殊减法的组合映射（用于处理 IV, IX, XL, XC, CD, CM）
    mapping(bytes => uint256) private subtractionValues;

    constructor() {
        // 初始化单个字符映射
        romanValues['I'] = 1;
        romanValues['V'] = 5;
        romanValues['X'] = 10;
        romanValues['L'] = 50;
        romanValues['C'] = 100;
        romanValues['D'] = 500;
        romanValues['M'] = 1000;

        // 初始化减法组合映射
        subtractionValues["IV"] = 4;
        subtractionValues["IX"] = 9;
        subtractionValues["XL"] = 40;
        subtractionValues["XC"] = 90;
        subtractionValues["CD"] = 400;
        subtractionValues["CM"] = 900;
    }

    /**
     * @dev 将罗马数字转换为整数（标准算法）
     * @param roman 罗马数字字符串（大写）
     * @return 对应的整数值
     */
    function romanToInt(string memory roman) public view returns (uint256) {
        bytes memory romanBytes = bytes(roman);
        require(romanBytes.length > 0, "Roman numeral cannot be empty");

        // 2942  MMCMXLII
        uint256 total = 0;

        for (uint256 i = 0; i < romanBytes.length; i++) {
            // 获取当前字符的值
            uint256 currentValue = getRomanValue(romanBytes[i]);

            // 如果不是最后一个字符，检查是否需要减法
            if (i < romanBytes.length - 1) {
                uint256 nextValue = getRomanValue(romanBytes[i + 1]);

                //MMCMXLII  8
                // 如果当前值小于下一个值，使用减法规则
                if (currentValue < nextValue) {
                    // 验证是否为有效的减法组合
                    require(isValidSubtraction(romanBytes[i], romanBytes[i + 1]),
                        "Invalid Roman numeral subtraction");

                    total += (nextValue - currentValue);
                    i++; // 跳过下一个字符，因为已经处理了
                } else {
                    total += currentValue;
                }
            } else {
                total += currentValue;
            }
        }

        // 验证结果是否在有效范围内（1-3999）
        require(total > 0 && total < 4000, "Result out of valid range (1-3999)");

        return total;
    }


    /**
     * @dev 获取单个罗马字符对应的数值
     * @param romanChar 罗马字符
     * @return 对应的整数值
     */
    function getRomanValue(bytes1 romanChar) public view returns (uint256) {
        uint256 value = romanValues[romanChar];
        require(value > 0, string(abi.encodePacked("Invalid Roman character: ", romanChar)));
        return value;
    }

    /**
     * @dev 验证是否为有效的减法组合
     * @param leftChar 左侧字符
     * @param rightChar 右侧字符
     * @return 是否为有效的减法组合
     */
    function isValidSubtraction(bytes1 leftChar, bytes1 rightChar) public pure returns (bool) {
        // 有效的减法组合只有6种：IV, IX, XL, XC, CD, CM
        return (leftChar == 'I' && (rightChar == 'V' || rightChar == 'X')) ||
            (leftChar == 'X' && (rightChar == 'L' || rightChar == 'C')) ||
            (leftChar == 'C' && (rightChar == 'D' || rightChar == 'M'));
    }

}