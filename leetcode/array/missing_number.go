package array

// 268. Missing Number
// https://leetcode.com/problems/missing-number
//
// O(n) time | O(1) space
func missingNumber(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return (len(nums) * (len(nums) + 1) / 2) - total
}
