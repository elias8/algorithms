package array

// 283. Move Zeroes
// https://leetcode.com/problems/move-zeroes/
//
// O(n) time | O(1) space
func moveZeroes(nums []int) {
	if len(nums) < 1 {
		return
	}
	left, right := 0, 1
	for right < len(nums) {
		if nums[left] == 0 && nums[right] != 0 {
			nums[left] = nums[right]
			nums[right] = 0
			right++
			left++
		}

		if right < len(nums) && nums[left] != 0 {
			right++
			left++
		} else if right < len(nums) && nums[right] == 0 {
			right++
		}
	}
}

// Solution 2
// O(n) time | O(1) space
//func moveZeroes(nums []int) {
//	snowBallSize := 0
//	for i := 0; i < len(nums); i++ {
//		if nums[i] == 0 {
//			snowBallSize++
//		} else if snowBallSize > 0 {
//			nums[i-snowBallSize] = nums[i]
//			nums[i] = 0
//		}
//	}
//}
