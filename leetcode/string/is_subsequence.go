package string

// 392. Is Subsequence
// https://leetcode.com/problems/is-subsequenc
//
// O(n) time | O(1) space
func isSubsequence(s string, t string) bool {
	j := 0
	for i := 0; i < len(t) && j < len(s); i++ {
		if s[j] == t[i] {
			j++
		}
	}
	return j == len(s)
}
