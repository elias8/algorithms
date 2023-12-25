package tree

// 100. Same Tree
// https://leetcode.com/problems/same-tree/
//
// O(n) time, where n is the smallest # of nodes among p & q
// O(h) space, where h is the height of the shorter tree
func isSameTree(p *Node, q *Node) bool {
	if p != nil && q != nil && p.Val == q.Val {
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}
	return p == nil && q == nil
}
