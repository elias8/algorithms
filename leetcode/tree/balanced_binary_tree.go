package tree

// 110. Balanced Binary Tree
// https://leetcode.com/problems/balanced-binary-tree
//
// O(n) time | O(n) space
func isBalanced(root *Node) bool {
	balanced, _ := checkBalanced(root)
	return balanced
}

func checkBalanced(root *Node) (bool, int) {
	if root == nil {
		return true, 0
	}
	leftBalanced, leftDepth := checkBalanced(root.Left)
	rightBalanced, rightDepth := checkBalanced(root.Right)
	diff := leftDepth - rightDepth
	if diff < 0 {
		diff = -diff
	}
	return leftBalanced && rightBalanced && diff <= 1, max(leftDepth, rightDepth) + 1
}
