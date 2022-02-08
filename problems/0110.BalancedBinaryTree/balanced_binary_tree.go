package problem110

import "github.com/deyuro/go-leetcode/structures"

func isBalanced(root *structures.TreeNode) bool {
	d := depth(root)
	if  d == -1{
		return false
	} else {
		return true
	}
}

func depth(root *structures.TreeNode) int {
	if root == nil {
		return 0
	}

	l := depth(root.Left)  // going left till the deadend
	r := depth(root.Right) // same for right

	// in has no balance anywhere return -1
	if l == -1 || r == -1 {
		return -1
	}

	diff := abs(l-r)
	if diff <= 1 {
		return max(l, r) + 1 // depth of longest + 1 for counting
	} else {
		return -1 // diff more than 1 tree not balances
	}

}

// absolute value of diff. Diff can not have minus sign, not matter which subtree is longer
func abs(val int) int {
	if val < 0 {
		return val * -1
	} else {
		return val
	}
}


func max(a, b int) int {
	if a > b {
		return a
	}else {
		return b
	}
}
