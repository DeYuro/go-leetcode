package problem141

import "github.com/deyuro/go-leetcode/structures"

func hasCycle(head *structures.ListNode) bool {
	if head == nil {
		return false
	}

	for s,f := head, head.Next; f != nil && f.Next != nil; s,f = s.Next,f.Next.Next {
		if s == f {
			return true
		}
	}

	return false
}

func hasCycle2(head *structures.ListNode) bool {
	m := make(map[*structures.ListNode]struct{})
	for head != nil {
		if _, ok := m[head]; ok {
			return true
		}
		m[head] = struct{}{}
		head = head.Next
	}
	return false
}