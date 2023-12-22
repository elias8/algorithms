package array

// 219. Contains Duplicate II
// https://leetcode.com/problems/contains-duplicate-ii
//
// O(n) time | O(n) space
func containsNearbyDuplicate(nums []int, k int) bool {
	window := make(map[int]bool, len(nums))
	for i, num := range nums {
		if _, found := window[num]; found {
			return true
		}
		window[num] = true
		if len(window) > k {
			delete(window, nums[i-k])
		}
	}
	return false
}
