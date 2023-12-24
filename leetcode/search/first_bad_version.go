package search

// isBadVersion is a mock function
func isBadVersion(n int) bool {
	return n%3 == 0
}

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
// 278. First Bad Version
// https://leetcode.com/problems/first-bad-version/
//
// O(log n) time | O(1) space
func firstBadVersion(n int) int {
	left, right := 0, n
	for left <= right {
		mid := (left + right) >> 1
		if isBadVersion(mid) && !isBadVersion(mid-1) {
			return mid
		} else if isBadVersion(mid - 1) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return 0
}
