package array

import "math"

// 202. Happy Number
// https://leetcode.com/problems/happy-number
//
// O(n) time | O(1) space
func isHappy(n int) bool {
	sum := 0
	num := n
	for num > 0 {
		sum += int(math.Pow(float64(num%10), 2))
		num = num / 10
		if num < 10 {
			sum += int(math.Pow(float64(num), 2))
			num = sum
			if sum == 1 {
				return true
			}
			sum = 0
		}
	}
	return false
}
