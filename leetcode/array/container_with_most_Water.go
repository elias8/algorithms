package array

// 11. Container With Most Water
// https://leetcode.com/problems/container-with-most-water
//
// O(n) time | O(1) space
func maxArea(height []int) int {
	maxArea, left, right := 0, 0, len(height)-1
	for left < right {
		distance := right - left
		minHeight := min(height[left], height[right])
		area := distance * minHeight
		maxArea = max(maxArea, area)

		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}

	return maxArea
}
