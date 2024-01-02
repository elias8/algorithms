package tree

// 589. N-ary Tree Preorder Traversal
// https://leetcode.com/problems/n-ary-tree-preorder-traversal/
//
// O(n) time || O(n) space
func preorder(root *NAryNode) []int {
	if root == nil {
		return nil
	}
	var result []int
	preorderHelper(root, &result)
	return result
}

func preorderHelper(root *NAryNode, result *[]int) {
	*result = append(*result, root.Val)
	for _, child := range root.Children {
		preorderHelper(child, result)
	}
}
