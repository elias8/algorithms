package string

// 2414. Length of the Longest Alphabetical Continuous Substring
// https://leetcode.com/problems/length-of-the-longest-alphabetical-continuous-substring
//
// O(n) time || O(1) space
func longestContinuousSubstring(s string) int {
	windowStart, maxLength := 0, 0
	for windowEnd := 0; windowEnd < len(s); windowEnd++ {
		if windowEnd > 0 && s[windowEnd]-s[windowEnd-1] != 1 {
			windowStart = windowEnd
		}
		windowSize := windowEnd - windowStart + 1
		if windowSize > maxLength {
			maxLength = windowSize
		}
	}
	return maxLength
}
