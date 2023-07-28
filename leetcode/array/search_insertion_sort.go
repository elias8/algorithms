package array

import "math"

// 35. Search Insert Position
// https://leetcode.com/problems/search-insert-position/
//
// O(log(n)) time | O(1) space
func searchInsert(nums []int, target int) int {
	left, right, mid := 0, len(nums)-1, 0
	for left <= right {
		mid = int(math.Trunc(float64((left + right) / 2)))
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left++
		} else {
			right--
		}
	}

	if target > nums[mid] {
		return mid + 1
	}
	return mid
}
