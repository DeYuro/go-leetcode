package problem142

import "github.com/deyuro/go-leetcode/structures"

func detectCycle(head *structures.ListNode) *structures.ListNode {
	m := make(map[*structures.ListNode]*structures.ListNode)

	for head != nil {
		if node,ok := m[head]; ok {
			return node
		}

		m[head] = head

		head = head.Next
	}

	return nil
}