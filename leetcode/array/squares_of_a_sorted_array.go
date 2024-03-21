package array

// 977. Squares of a Sorted Array
// https://leetcode.com/problems/squares-of-a-sorted-array
//
// O(n) time || O(n) space
func sortedSquares(nums []int) []int {
	result := make([]int, len(nums))
	lIndex, rIndex, pos := 0, len(nums)-1, len(nums)-1
	for lIndex <= rIndex {
		left := nums[lIndex] * nums[lIndex]
		right := nums[rIndex] * nums[rIndex]
		if left < right {
			result[pos] = right
			rIndex--
		} else {
			result[pos] = left
			lIndex++
		}
		pos--
	}
	return result
}
