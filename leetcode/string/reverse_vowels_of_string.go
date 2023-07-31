package string

import "unicode"

// 345. Reverse Vowels of a String
// https://leetcode.com/problems/reverse-vowels-of-a-string/
//
// O(n) time | O(n) space
func reverseVowels(s string) string {
	left, right := 0, len(s)-1
	text := make([]byte, len(s))
	for left <= right {
		if isVowel(s[left]) && isVowel(s[right]) {
			text[right] = s[left]
			text[left] = s[right]
			right--
			left++
		} else {
			if !isVowel(s[left]) {
				text[left] = s[left]
				left++
			}
			if !isVowel(s[right]) {
				text[right] = s[right]
				right--
			}
		}
	}
	return string(text)
}

func isVowel(char byte) bool {
	ch := unicode.ToLower(rune(char))
	return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u'
}
