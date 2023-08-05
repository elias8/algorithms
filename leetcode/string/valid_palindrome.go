package string

import "unicode"

// 125. Valid Palindrome
// https://leetcode.com/problems/valid-palindrome
//
// O(n) time | O(1) space
func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left <= right {
		if !isAlphaNumeric(s[left]) {
			left++
			continue
		}
		if !isAlphaNumeric(s[right]) {
			right--
			continue
		}

		if unicode.ToLower(rune(s[left])) != unicode.ToLower(rune(s[right])) {
			return false
		}
		if left == right {
			return true
		}
		right--
		left++
	}
	return true
}

func isAlphaNumeric(char byte) bool {
	return ('a' <= char && char <= 'z') || ('A' <= char && char <= 'Z') || ('0' <= char && char <= '9')
}
