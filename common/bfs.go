package main

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