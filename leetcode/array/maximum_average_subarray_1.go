package array

import "math"

// 643. Maximum Average Subarray I
// https://leetcode.com/problems/maximum-average-subarray-i
//
// O(n) time || O(1) space
func findMaxAverage(nums []int, k int) float64 {
	windowSum, maxSum := 0, math.MinInt
	for i := 0; i < len(nums); i++ {
		windowSum += nums[i]
		if i >= k-1 {
			if windowSum > maxSum {
				maxSum = windowSum
			}
			windowSum -= nums[i-k+1]
		}
	}

	return float64(maxSum) / float64(k)
}
