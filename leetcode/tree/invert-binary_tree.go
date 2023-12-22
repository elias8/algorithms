package tree

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// 226. Invert Binary Tree
// https://leetcode.com/problems/invert-binary-tree
//
// O(n) time | O(1) space
func invertTree(root *Node) *Node {
	stack := []*Node{root}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if current != nil {
			current.Left, current.Right = current.Right, current.Left
			stack = append(stack, current.Left, current.Right)
		}
	}

	return root
}

// Recursive solution
// O(n) time | O(d) space
// func invertTree(root *TreeNode) *TreeNode {
//     if root == nil {
//         return nil
//     }
//     root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
//     return root
// }
