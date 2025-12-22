// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract BinarySearch {

    /**
     * @dev 在有序数组中查找目标值（二分查找）
     * @param arr 已排序的升序数组
     * @param target 目标值
     * @return 目标值的索引，如果未找到则返回 -1
     */
    function binarySearch(uint256[] memory arr, uint256 target)
    public
    pure
    returns (int256)
    {
        if (arr.length == 0) {
            return -1;
        }

        uint256 left = 0;
        uint256 right = arr.length - 1;

        while (left <= right) {
            // 防止溢出：使用 left + (right - left) / 2
            uint256 mid = left + (right - left) / 2;

            if (arr[mid] == target) {
                return int256(mid);
            } else if (arr[mid] < target) {
                left = mid + 1;
            } else {
                // 防止下溢
                if (mid == 0) {
                    break;
                }
                right = mid - 1;
            }
        }

        return -1;
    }

}