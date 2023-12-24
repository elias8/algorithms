package array

// 35. Search Insert Position
// https://leetcode.com/problems/search-insert-position/
//
// O(log(n)) time | O(1) space
func searchInsert(nums []int, target int) int {
	left, mid, right := 0, 0, len(nums)-1
	for left <= right {
		mid = (left + right) >> 1
		if nums[mid] == target {
			return mid
		} else if target > nums[mid] {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid - 1
		}
	}

	if nums[mid] < target {
		return mid + 1
	}
	return mid
}
