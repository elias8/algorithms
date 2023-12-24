package array

func canConstruct(ransomNote string, magazine string) bool {
	letters := make([]int, 26)
	for _, char := range magazine {
		letters[char-'a']++
	}

	for _, char := range ransomNote {
		if letters[char-'a'] > 0 {
			letters[char-'a']--
		} else {
			return false
		}
	}
	return true
}
