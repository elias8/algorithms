package tree

import (
	"strconv"
	"strings"
)

// 606. Construct String from Binary Tree
// https://leetcode.com/problems/construct-string-from-binary-tree
//
// O(n) time || O(n) space
func tree2str(root *Node) string {
	builder := strings.Builder{}
	buildString(root, &builder)
	return builder.String()
}

func buildString(root *Node, builder *strings.Builder) {
	if root != nil {
		builder.WriteString(strconv.Itoa(root.Val))
		if root.Left != nil || root.Right != nil {
			builder.WriteByte('(')
			buildString(root.Left, builder)
			builder.WriteByte(')')
		}

		if root.Right != nil {
			builder.WriteByte('(')
			buildString(root.Right, builder)
			builder.WriteByte(')')
		}
	}
}
