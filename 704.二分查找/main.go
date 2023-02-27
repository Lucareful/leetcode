package main

import (
	"math"
)

func search(nums []int, target int) int {
	// 有序数组
	// 下标从 0 开始
	length := len(nums) - 1

	left := 0
	right := length

	for left <= right {
		mid := int(math.Ceil(float64((left + right) / 2)))
		switch {
		case target == nums[left]:
			return left
		case target == nums[right]:
			return right
		case target == nums[mid]:
			return mid
		case target > nums[mid] && mid != left:
			left = mid
		case target < nums[mid] && mid != right:
			right = mid
		default:
			return -1
		}
	}

	return -1
}

func BinarySearch(nums []int, target int, left int, right int, mid int) int {
	mid = int(math.Ceil(float64((left + right) / 2)))

	switch {
	case target == nums[left]:
		return left
	case target == nums[right]:
		return right
	case target == nums[mid]:
		return mid
	case target > nums[mid] && mid != left:
		left = mid
		return BinarySearch(nums, target, left, right, mid)
	case target < nums[mid] && mid != right:
		right = mid
		return BinarySearch(nums, target, left, right, mid)
	default:
		return -1
	}
}

// runtime:32 ms
// memory:6.5 MB
