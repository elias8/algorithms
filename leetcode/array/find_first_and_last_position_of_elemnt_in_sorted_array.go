package array

// 34. Find First and Last Position of Element in Sorted Array
// https://leetcode.com/problems/contains-duplicate-ii
//
// O(log n) time | O(1) space
func searchRange(nums []int, target int) []int {
	find := func(findLeft bool) int {
		position, left, right := -1, 0, len(nums)-1
		for left <= right {
			mid := (left + right) >> 1
			if nums[mid] == target {
				position = mid
				if findLeft {
					right = mid - 1
				} else {
					left = mid + 1
				}
			} else if nums[mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		return position
	}

	left, right := find(true), find(false)
	return []int{left, right}
}
