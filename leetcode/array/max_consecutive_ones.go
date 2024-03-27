package array

// 485. Max Consecutive Ones
// https://leetcode.com/problems/max-consecutive-ones/
//
// O(n) time || O(1) space
func findMaxConsecutiveOnes(nums []int) int {
	windowStart, maxConsecutive := 0, 0
	for windowEnd := 0; windowEnd < len(nums); windowEnd++ {
		if nums[windowEnd] == 0 {
			windowStart = windowEnd + 1
		}
		windowSize := windowEnd - windowStart + 1
		if windowSize > maxConsecutive {
			maxConsecutive = windowSize
		}
	}

	return maxConsecutive
}
