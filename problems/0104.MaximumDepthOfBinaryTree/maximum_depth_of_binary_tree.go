package problem104

import "github.com/deyuro/go-leetcode/structures"

func maxDepth(root *structures.TreeNode) int {
	return depth(root)
}

func depth(root *structures.TreeNode) int {
	if root == nil {
		return 0
	}

	l := depth(root.Left)
	r := depth(root.Right)


	return max(l,r) + 1
}

func max(a,b int) int {
	if a > b {
		return a
	}

	return b
}
