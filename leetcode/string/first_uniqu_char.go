package string

// 387. First Unique Character in a String
// https://leetcode.com/problems/first-unique-character-in-a-string/
//
// O(n) time | O(1) space
func firstUniqChar(s string) int {
	uniqueChars := make([]int, 26)
	for i := 0; i < len(s); i++ {
		uniqueChars[s[i]-'a']++
	}

	for i := 0; i < len(s); i++ {
		if uniqueChars[s[i]-'a'] == 1 {
			return i
		}
	}

	return -1
}
