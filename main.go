package main

import "math"

func main() {
	target := 14
	isPerfectSquare(target)
}

func isPerfectSquare(num int) bool {
	left := 0
	right := num

	for left < right {
		mid := int(math.Ceil(float64((left + right) / 2)))

		if mid*mid == num || (mid-1)*(mid-1) == num {
			return true
		}
		if num > mid*mid {
			left = mid
		} else {
			right = mid
		}
		if mid == left {
			return false
		}
	}

	return false
}
