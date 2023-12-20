package array

// 217. Contains Duplicate
// https://leetcode.com/problems/contains-duplicate
//
// O(n) time | O(n) space
func containsDuplicate(nums []int) bool {
	numsCount := map[int]int{}
	for _, num := range nums {
		if _, found := numsCount[num]; found {
			return true
		}
		numsCount[num]++
	}
	return false
}
