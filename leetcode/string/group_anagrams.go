package string

// 49. Group Anagrams
// https://leetcode.com/problems/group-anagrams/
//
// O(n^2) time | O(n) space
func groupAnagrams(strs []string) [][]string {
	anagrams := [][]string{}
	words := map[string]bool{}
	for _, word := range strs {
		words[word] = true
	}

	for i := 0; i < len(strs); i++ {
		word := strs[i]
		if _, found := words[word]; found {
			anagram := []string{word}
			for j := i + 1; j < len(strs); j++ {
				other := strs[j]
				if isAnagram(word, other) {
					anagram = append(anagram, other)
					delete(words, other)
				}
			}
			anagrams = append(anagrams, anagram)
		}
	}

	return anagrams
}
