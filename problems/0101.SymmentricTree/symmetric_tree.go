package problem100

import "github.com/deyuro/go-leetcode/structures"

func isSymmetric(root *structures.TreeNode) bool {
	return isSymmetricTree(root.Left, root.Right)
}

func isSymmetricTree(left, right *structures.TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left != nil && right == nil {
		return false
	}
	if left == nil && right != nil {
		return false
	}

	if left.Val == right.Val {
		return isSymmetricTree(left.Left, right.Right) && isSymmetricTree(left.Right, right.Left)
	}

	return false
}