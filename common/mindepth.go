package main

func minDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	left, right := minDepth(root.Left), minDepth(root.Right)
	if left == 0 || right == 0 {
		return left + right + 1
	}
	return min(left, right) + 1

}

func min(a,b int) int {
	if a < b {
		return a
	}

	return b
}

import "container/list"

type state struct {
	node *TreeNode
	lvl  int
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := list.New()
	queue.PushBack(&state{root, 1})

	for queue.Len() > 0 {
		v := queue.Front()
		q := v.Value.(*state)
		p, lvl := q.node, q.lvl
		queue.Remove(v)
		if p.Left == nil && p.Right == nil {
			return lvl
		}
		if p.Left != nil {
			queue.PushBack(&state{p.Left, lvl + 1})
		}
		if p.Right != nil {
			queue.PushBack(&state{p.Right, lvl + 1})
		}
	}
	return 0
}