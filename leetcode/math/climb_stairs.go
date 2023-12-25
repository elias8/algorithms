package math

// 70. Climbing Stairs
// https://leetcode.com/problems/climbing-stairs/
//
// O(n) time | O(1) space
func climbStairs(n int) int {
	a, b := 1, 1
	for i := 0; i < n-1; i++ {
		a, b = b, a+b
	}
	return b
}
