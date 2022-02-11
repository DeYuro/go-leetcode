package problem113

import "github.com/deyuro/go-leetcode/structures"

func pathSum(root *structures.TreeNode, targetSum int) [][]int {
	result := [][]int{}
	vals := []int{}
	pathSumHelper(root, targetSum, vals, &result)
	return result
}


func pathSumHelper(root *structures.TreeNode, targetSum int, vals []int,  res *[][]int) {

	if root == nil {
		return
	}
	vals = append(vals, root.Val)

	if root.Left == nil && root.Right == nil {
		if targetSum - root.Val == 0 {
			*res = append(*res, append([]int(nil),vals...))
		}
		return
	}

	pathSumHelper(root.Left, targetSum - root.Val, vals, res)
	pathSumHelper(root.Right, targetSum - root.Val, vals, res)
}
