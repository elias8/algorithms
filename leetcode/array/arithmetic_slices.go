package array

// 413. Arithmetic Slices
// https://leetcode.com/problems/arithmetic-slices
//
// O(n) time || O(1) space
func numberOfArithmeticSlices(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	windowStart, diff, count := 0, nums[0]-nums[1], 0
	for windowEnd := 2; windowEnd < len(nums); windowEnd++ {
		currentDiff := nums[windowEnd-1] - nums[windowEnd]
		if currentDiff != diff {
			count += countSlices(windowEnd - windowStart)
			windowStart = windowEnd - 1
			diff = currentDiff
		}
	}

	count += countSlices(len(nums) - windowStart)
	return count
}

func countSlices(length int) int {
	if length < 3 {
		return 0
	}
	return (length - 1) * (length - 2) / 2
}
