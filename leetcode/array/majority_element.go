package array

// 169. Majority Element
// https://leetcode.com/problems/majority-element
//
// O(n) time | O(n) time
func majorityElement(nums []int) int {
	elements := map[int]int{}
	majorityNum, majorityElement := 0, 0
	for i := 0; i < len(nums); i++ {
		elements[nums[i]]++
		if elements[nums[i]] > len(nums)/2 && elements[nums[i]] > majorityElement {
			majorityElement = elements[nums[i]]
			majorityNum = nums[i]
		}
	}
	return majorityNum
}
