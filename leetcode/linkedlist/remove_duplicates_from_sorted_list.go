package linkedlist

// 83. Remove Duplicates from Sorted List
// https://leetcode.com/problems/remove-duplicates-from-sorted-list/
//
// O(n) time || O(1) space
func deleteDuplicates(head *ListNode) *ListNode {
	current := head
	for current != nil && current.Next != nil {
		if current.Val == current.Next.Val {
			if current.Next.Next != nil {
				current.Next = current.Next.Next
			} else {
				current.Next = nil
			}
			continue
		}
		current = current.Next
	}
	return head
}
