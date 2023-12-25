package array

// 268. Missing Number
// https://leetcode.com/problems/missing-number
//
// O(n) time | O(1) space
func missingNumber(nums []int) int {
	missingNumber := len(nums) * (len(nums) + 1) / 2
	for _, num := range nums {
		missingNumber -= num
	}
	return missingNumber
}
