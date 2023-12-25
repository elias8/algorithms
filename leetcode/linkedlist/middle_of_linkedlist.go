package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

// 876. Middle of the Linked List
// https://leetcode.com/problems/middle-of-the-linked-list/
//
// O(n) time | O(n) space
func middleNode(head *ListNode) *ListNode {
	indexes, count := make(map[int]*ListNode), 0
	for head != nil {
		indexes[count] = head
		head = head.Next
		count++
	}
	return indexes[count/2]
}
