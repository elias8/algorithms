package string

// 20. Valid Parentheses
// https://leetcode.com/problems/valid-parentheses/
//
// O(n) time | O(1) space
func isValid(s string) bool {
	var stack []rune
	for _, char := range s {
		if len(stack) > 0 {
			if isClosing(stack[len(stack)-1], char) {
				stack = stack[:len(stack)-1]
				continue
			}
		}
		stack = append(stack, char)
	}
	return len(stack) == 0
}

func isClosing(prev, curr rune) bool {
	return prev == '(' && curr == ')' ||
		prev == '[' && curr == ']' ||
		prev == '{' && curr == '}'
}
