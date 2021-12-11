package problem100

import (
	"github.com/deyuro/go-leetcode/structures"
	"reflect"
)

func isSameTree(p *structures.TreeNode, q *structures.TreeNode) bool {
	return reflect.DeepEqual(postorderTraversal(p), postorderTraversal(q))
}

func postorderTraversal(root *structures.TreeNode) []int {

	res := []int{} // init for not nil

	if root == nil {
		return res
	}

	postorder(root, &res)

	return res
}

func postorder(root *structures.TreeNode, res *[]int) {

	if root.Left != nil {
		postorder(root.Left, res)
	}

	if root.Right != nil {
		postorder(root.Right, res)
	}

	*res = append(*res, root.Val)
}