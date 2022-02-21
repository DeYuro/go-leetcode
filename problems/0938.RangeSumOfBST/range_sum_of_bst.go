package problem938

import "github.com/deyuro/go-leetcode/structures"

func rangeSumBST(root *structures.TreeNode, low int, high int) int {
	res := 0
	sum(root, low,high,&res)
	return res
}

func sum(root *structures.TreeNode, low, high int, res *int) {
	if root == nil {
		return
	}

	if low <= root.Val && root.Val <= high {
		*res += root.Val
	}

	sum(root.Left, low, high, res)
	sum(root.Right, low, high, res)
}