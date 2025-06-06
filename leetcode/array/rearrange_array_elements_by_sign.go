package array

// 2149. Rearrange Array Elements by Sign
// https://leetcode.com/problems/rearrange-array-elements-by-sign
//
// O(n) time | O(n) space
func rearrangeArray(nums []int) []int {
	rearranged := make([]int, len(nums))
	positive, negative := 0, 1

	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			rearranged[negative] = nums[i]
			negative += 2
		} else {
			rearranged[positive] = nums[i]
			positive += 2
		}
	}

	return rearranged
}
