package problem145

import "github.com/deyuro/go-leetcode/structures"

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