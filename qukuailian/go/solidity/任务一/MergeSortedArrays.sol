// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

//将两个有序数组合并为一个有序数组。

contract MergeSortedArrays {
    //将两个有序数组合并为一个有序数组。
    //将两个数组合并成一个数据
    //去重
    //排序


    function mergeSortAndDeduplicate(
        uint256[] memory arr1,
        uint256[] memory arr2
    ) public pure returns (uint256[] memory) {
        // 合并两个数组
        uint256[] memory merged = mergeArrays(arr1, arr2);
        //排序
        uint256[] memory mergedList = sort(merged);
        //去重




        // 去重
        uint256[] memory deduplicated = removeDuplicatesFromSorted(merged);

        // 排序（由于输入已有序且合并时保持有序，所以这里主要是保证结果有序）
        // 实际上，合并有序数组的结果已经是有序的，去重后仍然有序
        return deduplicated;
    }

    /**
     * @dev 对数组进行排序（冒泡排序，适合小数组）
     * @param arr 要排序的数组
     * @return 排序后的数组
     */
    function sort(uint256[] memory arr) public pure returns (uint256[] memory) {
        uint256[] memory sorted = new uint256[](arr.length);

        // 复制原数组
        for (uint256 i = 0; i < arr.length; i++) {
            sorted[i] = arr[i];
        }

        // 冒泡排序
        for (uint256 i = 0; i < sorted.length; i++) {
            for (uint256 j = i + 1; j < sorted.length; j++) {
                if (sorted[i] > sorted[j]) {
                    // 交换
                    (sorted[i], sorted[j]) = (sorted[j], sorted[i]);
                }
            }
        }

        return sorted;
    }

    /**
     * @dev 从已排序数组中移除重复元素
     * @param sortedArr 已排序的数组
     * @return 去重后的数组
     */
    function removeDuplicatesFromSorted(uint256[] memory sortedArr)
    public
    pure
    returns (uint256[] memory)
    {
        if (sortedArr.length == 0) {
            return new uint256[](0);
        }

        // 计算不重复元素的数量
        uint256 uniqueCount = 1;
        for (uint256 i = 1; i < sortedArr.length; i++) {
            if (sortedArr[i] != sortedArr[i - 1]) {
                uniqueCount++;
            }
        }

        // 创建结果数组
        uint256[] memory result = new uint256[](uniqueCount);
        result[0] = sortedArr[0];

        // 填充结果数组
        uint256 resultIndex = 1;
        for (uint256 i = 1; i < sortedArr.length; i++) {
            if (sortedArr[i] != sortedArr[i - 1]) {
                result[resultIndex] = sortedArr[i];
                resultIndex++;
            }
        }

        return result;
    }


    /**
     * @dev 合并两个数组（适用于任意数组，不一定有序）
     * @param arr1 第一个数组
     * @param arr2 第二个数组
     * @return 合并后的数组
     */
    function mergeArrays(
        uint256[] memory arr1,
        uint256[] memory arr2
    ) public pure returns (uint256[] memory) {
        uint256[] memory merged = new uint256[](arr1.length + arr2.length);

        // 复制第一个数组
        for (uint256 i = 0; i < arr1.length; i++) {
            merged[i] = arr1[i];
        }

        // 复制第二个数组
        for (uint256 i = 0; i < arr2.length; i++) {
            merged[arr1.length + i] = arr2[i];
        }

        return merged;
    }
}