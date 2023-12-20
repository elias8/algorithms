package array

// 303. Range Sum Query - Immutable
// https://leetcode.com/problems/range-sum-query-immutable/

type NumArray struct {
	nums []int
}

// Constructor O(n) time | O(n) space
func Constructor(nums []int) NumArray {
	sums := make([]int, len(nums))
	sums[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sums[i] = sums[i-1] + nums[i]
	}
	return NumArray{sums}
}

// SumRange O(1) time | O(1) space
func (a *NumArray) SumRange(left int, right int) int {
	if left < 1 {
		return a.nums[right]
	}
	return a.nums[right] - a.nums[left-1]
}
