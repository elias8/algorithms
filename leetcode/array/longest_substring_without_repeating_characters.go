package array

// 3. Longest Substring Without Repeating Characters
// https://leetcode.com/problems/longest-substring-without-repeating-characters
//
// O(n) time || O(n) space
func lengthOfLongestSubstring(s string) int {
	windowStart, maxWindowDistance, charIndex := 0, 0, map[uint8]int{}
	for windowEnd := 0; windowEnd < len(s); windowEnd++ {
		char := s[windowEnd]
		if index, found := charIndex[char]; found && windowStart <= index {
			windowStart = index + 1
		}
		windowDistance := windowEnd - windowStart + 1
		if windowDistance > maxWindowDistance {
			maxWindowDistance = windowDistance
		}
		charIndex[char] = windowEnd
	}

	return maxWindowDistance
}
