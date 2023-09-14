package string

// 58. Length of Last Word
// https://leetcode.com/problems/length-of-last-word
//
// O(n) time | O(1) space
func lengthOfLastWord(s string) int {
	lastWordLength := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if lastWordLength > 0 {
				break
			}
			continue
		}
		lastWordLength++
	}
	return lastWordLength
}
