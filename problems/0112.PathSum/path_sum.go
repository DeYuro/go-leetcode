package problem112

import "github.com/deyuro/go-leetcode/structures"

func hasPathSum(root *structures.TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		if targetSum - root.Val == 0 {
			return true
		}
	}

	return hasPathSum(root.Left, targetSum - root.Val) || hasPathSum(root.Right, targetSum - root.Val)
}