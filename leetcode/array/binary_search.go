package array

// 704. Binary Search
// https://leetcode.com/problems/binary-search/
//
// O(log n) time | O(1) space
func search2(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) >> 1
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
