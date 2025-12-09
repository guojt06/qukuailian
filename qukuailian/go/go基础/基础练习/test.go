package main

import (
	"fmt"
	"sort"
	"strconv"
)

//  1.给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。
// 找出那个只出现了一次的元素。可以使用 `for` 循环遍历数组，结合 `if` 条件判断和 `map` 数据结构来解决，
// 例如通过 `map` 记录每个元素出现的次数，然后再遍历 `map` 找到出现次数为1的元素

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// 交换元素
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func findString(arr []int) {
	map1 := make(map[int]int)
	for _, num := range arr {
		map1[num] += 1
	}
	for key, value := range map1 {
		if value == 1 {
			fmt.Println("元素 " + strconv.Itoa(key) + " 出现了 " + strconv.Itoa(value) + " 次")
		}
		// fmt.Println("元素 " + strconv.Itoa(key) + " 出现了 " + strconv.Itoa(value) + " 次")
	}

}

// 判断一个整数是否是回文数
// 先将这个回文数转换为字符串，然后将字符串反转后与原字符串比较
func isPalindrome(s string) bool {
	var a string
	runes := []rune(s)
	for i := len(runes) - 1; i > 0; i-- {
		a = a + string(runes[i])
	}
	return a == s
}

//给定一个由整数组成的非空数组所表现的非负整数，在该数的基础上加一

func plusOne(digits []int) []int {
	for i := 0; i < len(digits); i++ {
		digits[i] = digits[i] + 1
	}
	return digits
}

// 4.- **[26. 删除有序数组中的重复项](https://leetcode.cn/problems/remove-duplicates-from-sorted-array/ "26. 删除有序数组中的重复项
// (https://leetcode.cn/problems/remove-duplicates-from-sorted-array/)")**：给你一个有序数组 `nums` ，请你原地删除重复出现的元素，
// 使每个元素只出现一次，返回删除后数组的新长度。不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成
// 。可以使用双指针法，一个慢指针 `i` 用于记录不重复元素的位置，一个快指针 `j` 用于遍历数组，当 `nums[i]` 与 `nums[j]` 不相等时，将 `nums[j]` 赋值给 `nums[i + 1]`，
// 并将 `i` 后移一位。
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i = i + 1
			nums[i] = nums[j]
		}
	}
	return i + 1
}

// 以数组 `intervals` 表示若干个区间的集合，其中单个区间为 `intervals[i] = [starti, endi]` 。
// 请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。可以先对区间数组按照区间的起始位置进行排序，
// 然后使用一个切片来存储合并后的区间，遍历排序后的区间数组，将当前区间与切片中最后一个区间进行比较，如果有重叠，则合并区间；如果没有重叠，则将当前区间添加到切片中。
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}
	// 先对区间数组按照区间的起始位置进行排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	// 初始化结果切片，先将第一个区间加入
	result := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		// 如果当前区间与切片中最后一个区间有重叠，则合并区间
		if intervals[i][0] <= result[len(result)-1][1] {
			result[len(result)-1][1] = max(intervals[i][1], result[len(result)-1][1])
		} else {
			// 如果没有重叠，则将当前区间添加到切片中
			result = append(result, intervals[i])
		}
	}
	return result
}

// 6.给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
// ，并返回他们的数组下标。你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
// 你可以按任意顺序返回答案。
func twoSum(nums []int, target int) []int {
	map1 := make(map[int]int)
	for i, num := range nums {
		if _, ok := map1[target-num]; ok {
			return []int{map1[target-num], i}
		}
		map1[num] = i
	}
	return nil
}

func main() {
	// arr := []int{1, 2, 3, 4, 5, 6, 1, 2, 4, 5, 6}
	// // fmt.Println("排序前:", arr)
	// // bubbleSort(arr)
	// // fmt.Println("排序后:", arr)
	// // findString(arr)
	// plusOne(arr)
	// fmt.Println(arr)

	// nums := []int{0, 0, 1, 1, 1, 2, 2, 2, 3, 3, 4}
	// ums := []int{0, 1, 2, 3, 4}
	// fmt.Println("原数组:", nums)
	// length := removeDuplicates(nums)
	// fmt.Println("新长度:", length)
	// fmt.Println("修改后数组:", nums[:length])
	// fmt.Println("原数组:", nums)

	ums := []int{0, 1, 2, 3, 4}
	// fmt.Println("原数组:", nums)
	length := twoSum(ums, 5)
	fmt.Println("新长度:", length)
	// fmt.Println("修改后数组:", nums[:length])

}
