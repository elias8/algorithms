package string

// 344. Reverse String
// https://leetcode.com/problems/reverse-string/
//
// O(n) time | O(1) space
func reverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		right--
		left++
	}
}
