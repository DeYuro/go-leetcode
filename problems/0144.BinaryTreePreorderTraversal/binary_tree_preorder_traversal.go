package problem144

import "github.com/deyuro/go-leetcode/structures"

func preorderTraversal(root *structures.TreeNode) []int {

	res := []int{} // init for not nil

	if root == nil {
		return res
	}

	preorder(root, &res)

	return res
}

func preorder(root *structures.TreeNode, res *[]int) {

	*res = append(*res, root.Val)

	if root.Left != nil {
		preorder(root.Left, res)
	}

	if root.Right != nil {
		preorder(root.Right, res)
	}
}