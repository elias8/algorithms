package tree

// 94. Binary Tree Inorder Traversal
// https://leetcode.com/problems/binary-tree-inorder-traversal/
//
// O(n) time || O(n) space
func inorderTraversal(root *Node) []int {
	if root == nil {
		return []int{}
	}
	var result = inorderTraversal(root.Left)
	result = append(result, root.Val)
	result = append(result, inorderTraversal(root.Right)...)
	return result
}
