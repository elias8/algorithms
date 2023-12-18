package array

// 1. Two Sum
// https://leetcode.com/problems/two-sum/
//
// O(n) time | O(n) space
func twoSum(nums []int, target int) []int {
	missingNumbers := map[int]int{}
	for index, num := range nums {
		missing := target - num
		if other, found := missingNumbers[missing]; found {
			return []int{index, other}
		}
		missingNumbers[num] = index
	}
	return []int{}
}
