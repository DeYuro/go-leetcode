package problem100

import (
	"github.com/deyuro/go-leetcode/structures"
)


func isSameTree(p *structures.TreeNode, q *structures.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p != nil && q == nil {
		return false
	}

	if p == nil && q != nil {
		return false
	}


	if  p.Val == q.Val {
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}

	return false
}