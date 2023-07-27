package array

// 27. Remove Element
// https://leetcode.com/problems/remove-element/
//
// O(n) time | O(1) space
func removeElement(nums []int, val int) int {
	var k int
	for _, num := range nums {
		if num != val {
			nums[k] = num
			k++
		}
	}
	return k
}
