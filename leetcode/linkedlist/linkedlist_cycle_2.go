package linkedlist

// 142. Linked List Cycle II
// https://leetcode.com/problems/linked-list-cycle-ii
//
// O(n) time || O(1) space
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			slow = head
			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}
			return slow
		}
	}
	return nil
}

// O(n) time || O(n) space
//func detectCycle(head *ListNode) *ListNode {
//	slow, fast := head, head
//	seen := make(map[*ListNode]bool)
//	for fast != nil && fast.Next != nil {
//		if seen[slow] {
//			return slow
//		}
//		seen[slow] = true
//		slow = slow.Next
//		fast = fast.Next.Next
//	}
//	return nil
//}
