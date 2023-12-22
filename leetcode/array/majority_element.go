package array

// 169. Majority Element
// https://leetcode.com/problems/majority-element
//
// O(n) time | O(1) time
func majorityElement(nums []int) int {
	var major, frequency int
	for _, num := range nums {
		if frequency == 0 {
			major = num
		}
		if num == major {
			frequency++
		} else {
			frequency--
		}
	}
	return major
}
