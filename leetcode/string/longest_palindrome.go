package string

// 409. Longest Palindrome
// https://leetcode.com/problems/longest-palindrome/
//
// O(n) time || O(1) space
func longestPalindrome(s string) int {
	letters := [52]int{}
	for _, char := range s {
		letters[indexOfChar(char)]++
	}

	pairs, hasOdd := 0, false
	for _, count := range letters {
		if count%2 == 0 {
			pairs += count
		} else if count > 0 {
			pairs += count - 1
			hasOdd = true
		}
	}

	if hasOdd {
		return pairs + 1
	}
	return pairs
}

func indexOfChar(char rune) int {
	if char >= 'A' && char <= 'Z' {
		return int(char - 'A')
	}
	return 26 + int(char-'a')
}
