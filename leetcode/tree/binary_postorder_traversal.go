package tree

// 145. Binary Tree Postorder Traversal
// https://leetcode.com/problems/binary-tree-postorder-traversal/
//
// O(n) time || O(n) space
func postorderTraversal(root *Node) []int {
	if root == nil {
		return []int{}
	}
	result := postorderTraversal(root.Left)
	result = append(result, postorderTraversal(root.Right)...)
	result = append(result, root.Val)
	return result
}
