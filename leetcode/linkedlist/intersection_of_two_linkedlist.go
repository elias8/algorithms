package linkedlist

import "math"

// 160. Intersection of Two Linked Lists
// https://leetcode.com/problems/intersection-of-two-linked-lists/
//
// O(m+n) time | O(1) space
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var lenA, lenB int
	for curr1, curr2 := headA, headB; curr1 != nil || curr2 != nil; {
		if curr1 != nil {
			curr1 = curr1.Next
			lenA++
		}
		if curr2 != nil {
			curr2 = curr2.Next
			lenB++
		}
	}

	currA, currB := headA, headB
	for i := 0; i < int(math.Abs(float64(lenA-lenB))); i++ {
		if lenA > lenB {
			currA = currA.Next
		} else {
			currB = currB.Next
		}
	}

	for currA != nil && currB != nil {
		if currA == currB {
			return currA
		}
		currA = currA.Next
		currB = currB.Next
	}
	return nil
}
