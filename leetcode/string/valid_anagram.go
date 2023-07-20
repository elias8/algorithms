package string

// 242. https://leetcode.com/problems/valid-anagram/submissions/
//
// O(n) time | O(c) space
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	alphabets := make([]int, 26)
	for i := 0; i < len(s); i++ {
		alphabets[s[i]-'a']++
		alphabets[t[i]-'a']--
	}

	for _, v := range alphabets {
		if v != 0 {
			return false
		}
	}

	return true
}
