package linkedlist

// 86. Partition List
// https://leetcode.com/problems/partition-list
//
// O(n) time | O(1) space
func partition(head *ListNode, x int) *ListNode {
	lessDummy, greaterDummy := &ListNode{}, &ListNode{}
	lessTail, greaterTail := lessDummy, greaterDummy

	curr := head
	for curr != nil {
		if curr.Val < x {
			lessTail.Next = curr
			lessTail = lessTail.Next
		} else {
			greaterTail.Next = curr
			greaterTail = greaterTail.Next
		}
		curr = curr.Next
	}

	lessTail.Next = greaterDummy.Next
	greaterTail.Next = nil
	return lessDummy.Next
}
