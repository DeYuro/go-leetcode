package problem94

import "github.com/deyuro/go-leetcode/structures"

func inTraversal(root *structures.TreeNode) []int {

	res := []int{} // init for not nil

	if root == nil {
		return res
	}

	inorder(root, &res)

	return res
}

func inorder(root *structures.TreeNode, res *[]int) {

	if root.Left != nil {
		inorder(root.Left, res)
	}

	*res = append(*res, root.Val)

	if root.Right != nil {
		inorder(root.Right, res)
	}
}