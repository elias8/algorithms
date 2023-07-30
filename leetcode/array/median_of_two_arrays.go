package array

// 4. Median of Two Sorted Arrays
// https://leetcode.com/problems/median-of-two-sorted-arrays/
//
// O(n) time | O(m+n) space
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var left, right int
	merged := make([]int, len(nums1)+len(nums2))
	for i := 0; i < len(merged); i++ {
		if left < len(nums1) && right < len(nums2) {
			if nums1[left] <= nums2[right] {
				merged[i] = nums1[left]
				left++
			} else {
				merged[i] = nums2[right]
				right++
			}
		} else if left < len(nums1) {
			merged[i] = nums1[left]
			left++
		} else if right < len(nums2) {
			merged[i] = nums2[right]
			right++
		}
	}

	if len(merged) == 1 {
		return float64(merged[0])
	}

	mid := len(merged) / 2
	if len(merged)%2 != 0 {
		return float64(merged[mid])
	} else {
		return float64(merged[mid-1]+merged[mid]) / 2
	}
}
