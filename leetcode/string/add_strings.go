package string

import "math"

// 415. Add Strings
// https://leetcode.com/problems/add-strings/
//
// O(n) time | O(n) space - where n is the length of the longer string
func addStrings(num1 string, num2 string) string {
	i, j := len(num1)-1, len(num2)-1
	longest := int(math.Max(float64(len(num1)), float64(len(num2))))
	answer, carry := make([]rune, longest+1), rune(0)
	for longest >= 0 {
		n1, n2 := rune(0), rune(0)
		if i >= 0 {
			n1 = rune(num1[i] - '0')
			i--
		}
		if j >= 0 {
			n2 = rune(num2[j] - '0')
			j--
		}

		sum := carry + n1 + n2
		answer[longest] = sum%10 + '0'
		if sum > 9 {
			carry = 1
		} else {
			carry = 0
		}
		longest--
		if longest == 0 && carry > 0 {
			answer[longest] = '1'
			return string(answer)
		}
	}

	return string(answer[1:])
}
