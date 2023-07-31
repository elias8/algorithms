package string

import "strconv"

// 412. Fizz Buzz
// https://leetcode.com/problems/fizz-buzz/solutions/
//
// O(n) time | O(n) space
func fizzBuzz(n int) []string {
	answer := make([]string, n)
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			answer[i-1] = "FizzBuzz"
		} else if i%3 == 0 {
			answer[i-1] = "Fizz"
		} else if i%5 == 0 {
			answer[i-1] = "Buzz"
		} else {
			answer[i-1] = strconv.Itoa(i)
		}
	}
	return answer
}
