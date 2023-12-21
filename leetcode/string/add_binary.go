package string

// 67. Add Binary
// https://leetcode.com/problems/add-binary/
//
// O(n) time | O(n) space
func addBinary(a string, b string) string {
	var carry uint8
	var result []rune
	i, j := len(a)-1, len(b)-1
	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry
		if i >= 0 {
			sum += a[i] - '0'
			i--
		}
		if j >= 0 {
			sum += b[j] - '0'
			j--
		}
		result = append([]rune{rune(sum%2) + '0'}, result...)
		carry = sum / 2
	}

	return string(result)
}
