/**
给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。

 如果数组中不存在目标值 target，返回 [-1, -1]。

 你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。



 示例 1：


输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]

 示例 2：


输入：nums = [5,7,7,8,8,10], target = 6
输出：[-1,-1]

 示例 3：


输入：nums = [], target = 0
输出：[-1,-1]



 提示：


 0 <= nums.length <= 10⁵
 -10⁹ <= nums[i] <= 10⁹
 nums 是一个非递减数组
 -10⁹ <= target <= 10⁹


 Related Topics 数组 二分查找 👍 2166 👎 0

*/

/* 题解思路：
此题要求 log（n）级别的复杂度程序，所以采用二分法，便可以将此题拆分为两个部分：
	1.找到左边界大于等于target的数组下标的值
	2.找到右边界小于等于target的数组下标的值

情况讨论：
	1.边界分析 如果 nums 为空，或者 target 小于 nums 的最小值，或 target 大于nums的最大值 返回 []int{-1,-1}

	2.target 在数组范围中，且数组中不存在target，例如数组{3,6,7},target为5，此时应该返回{-1, -1}

	3.target 在数组范围中，且数组中存在target，例如数组{3,6,7},target为6，此时应该返回{1, 1}


*/

package main

import (
	"math"
)

// leetcode submit region begin(Prohibit modification and deletion)
func searchRange(nums []int, target int) []int {
	var (
		left  = 0
		right = len(nums) - 1
	)
	if len(nums) == 0 || target < nums[left] || target > nums[right] {
		return []int{-1, -1}
	}

	leftBorder := findLeftBorder(nums, target)
	rightBorder := findRightBorder(nums, target)

	if leftBorder == -2 || rightBorder == -2 {
		return []int{-1, -1}
	}

	if rightBorder-leftBorder > 1 {
		return []int{leftBorder + 1, rightBorder - 1}
	}

	return []int{-1, -1}
}

func findLeftBorder(nums []int, target int) int {
	var (
		left       = 0
		right      = len(nums) - 1
		leftBorder = -2
	)
	for left <= right {
		middle := int(math.Ceil(float64((left + right) / 2)))

		if nums[middle] >= target {
			right = middle - 1
			leftBorder = right
		} else {
			left = middle + 1
		}
	}
	return leftBorder
}

func findRightBorder(nums []int, target int) int {
	var (
		left        = 0
		right       = len(nums) - 1
		rightBorder = -2
	)
	for left <= right {
		middle := int(math.Ceil(float64((left + right) / 2)))
		if nums[middle] > target {
			right = middle - 1
		} else {
			left = middle + 1
			rightBorder = left
		}
	}
	return rightBorder
}

// leetcode submit region end(Prohibit modification and deletion)
