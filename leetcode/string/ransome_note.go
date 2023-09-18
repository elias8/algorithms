package string

// 383. Ransom Note
// https://leetcode.com/problems/ransom-note
//
// O(n) time | O(1) space - where n is the length of the ransom note
func canConstruct(ransomNote string, magazine string) bool {
	alphabet := make([]int, 26)
	for _, char := range magazine {
		alphabet[char-'a']++
	}

	for _, char := range ransomNote {
		if alphabet[char-'a'] == 0 {
			return false
		}
		alphabet[char-'a']--
	}

	return true
}
