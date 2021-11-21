package problem144

func preorderTraversal(root *TreeNode) []int {

	var res []int

	if root == nil {
		return res
	}

	preorder(root, &res)

	return res
}

func preorder(root *TreeNode, res *[]int) {

	*res = append(*res, root.Val)

	if root.Left != nil {
		preorder(root.Left, res)
	}

	if root.Right != nil {
		preorder(root.Right, res)
	}
}