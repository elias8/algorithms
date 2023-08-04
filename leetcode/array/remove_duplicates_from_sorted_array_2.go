package array

// 80. Remove Duplicates from Sorted Array II
// https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii
//
// O(n) time | O(1) space
func removeDuplicates2(nums []int) int {
	k := 0
	for i := 0; i < len(nums); i++ {
		if i < 2 || nums[i] != nums[k-2] {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}
