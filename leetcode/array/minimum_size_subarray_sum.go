package array

import "math"

// 209. Minimum Size Subarray Sum
// https://leetcode.com/problems/minimum-size-subarray-sum/
//
// O(n) time || O(1) space
func minSubArrayLen(target int, nums []int) int {
	windowStart, windowEnd, windowSum, minWindowDistance := 0, 0, 0, math.MaxInt
	for windowStart < len(nums) {
		if windowSum < target && windowEnd < len(nums) {
			windowSum += nums[windowEnd]
			windowEnd++
		} else if windowSum >= target {
			windowSize := windowEnd - windowStart
			if windowSize < minWindowDistance {
				minWindowDistance = windowSize
			}
			windowSum -= nums[windowStart]
			windowStart++
		} else {
			break
		}
	}
	if minWindowDistance == math.MaxInt {
		return 0
	}
	return minWindowDistance
}

// O(n) time | O(1) space
//func minSubArrayLen(target int, nums []int) int {
//	left, right := 0, 0
//	minLen, sum := math.MaxInt, 0
//	for right < len(nums) || sum > target && left < len(nums) {
//		if sum >= target {
//			if right-left <= minLen {
//				minLen = right - left
//			}
//			sum -= nums[left]
//			left++
//		} else if right < len(nums) {
//			sum += nums[right]
//			right++
//		}
//	}
//
//	if minLen < math.MaxInt {
//		return minLen
//	}
//	return 0
//}
