package array

// 2460. Apply Operations to an Array
// https://leetcode.com/problems/apply-operations-to-an-array/
//
// O(n) time | O(1) space
func applyOperations(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			nums[i+1] = 0
			nums[i] *= 2
		}
	}

	var snowBallSize int
	for index, num := range nums {
		if num == 0 {
			snowBallSize++
		} else if snowBallSize > 0 {
			nums[index-snowBallSize] = num
			nums[index] = 0
		}
	}

	return nums
}
