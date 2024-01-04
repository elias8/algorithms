package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	num1, num2 := 0, 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			num1 = num1*10 + l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			num2 = num2*10 + l2.Val
			l2 = l2.Next
		}
	}

	head := &ListNode{}
	current := head
	sum := num1 + num2
	if sum == 0 {
		return &ListNode{Val: 0}
	}
	for sum > 0 {
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
		sum = sum / 10
	}
	return head.Next
}

func main() {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}

	addTwoNumbers(l1, l2)
}
