package linkedlist

// 141. Linked List Cycle
// https://leetcode.com/problems/linked-list-cycle
//
// O(n) time || O(1) space
func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}

	return false
}
