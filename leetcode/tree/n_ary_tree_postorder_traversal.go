package tree

type NAryNode struct {
	Val      int
	Children []*NAryNode
}

// 590. N-ary Tree Postorder Traversal
// https://leetcode.com/problems/n-ary-tree-postorder-traversal/
//
// O(n) time || O(n) space
func postorder(root *NAryNode) []int {
	if root == nil {
		return nil
	}
	var result []int
	postorderHelper(root, &result)
	return result
}

func postorderHelper(root *NAryNode, result *[]int) {
	for _, child := range root.Children {
		postorderHelper(child, result)
	}
	*result = append(*result, root.Val)
}
