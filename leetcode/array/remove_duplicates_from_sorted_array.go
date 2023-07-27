package array

// 26. Remove Duplicates from Sorted Array
// https://leetcode.com/problems/remove-duplicates-from-sorted-array/
//
// O(n) time | O(1) space
func removeDuplicates(nums []int) int {
	var i int
	for _, num := range nums {
		if num != nums[i] {
			i++
		}
		nums[i] = num
	}
	return i + 1
}
