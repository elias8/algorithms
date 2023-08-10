package math

// 9. Palindrome Number
// https://leetcode.com/problems/palindrome-number/
//
// O(n) time | O(1) space - where n is the number of digits in the input number
func isPalindrome(x int) bool {
	reversed, temp := 0, x
	for temp > 0 {
		digit := temp % 10
		reversed = reversed*10 + digit
		temp = temp / 10
	}
	return reversed == x
}
