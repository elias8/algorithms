package array

// 88. Merge Sorted Array
// https://leetcode.com/problems/merge-sorted-array/
//
// O(n) time | O(1) space
func merge(nums1 []int, m int, nums2 []int, n int) {
	left, right, index := m-1, n-1, m+n-1
	for right >= 0 {
		if left >= 0 && nums1[left] >= nums2[right] {
			nums1[index] = nums1[left]
			left--
		} else {
			nums1[index] = nums2[right]
			right--
		}
		index--
	}
}
