package problem21

import "github.com/go-leetcode/structures"

func mergeTwoLists(l1 *structures.ListNode,l2 *structures.ListNode) *structures.ListNode {
	result := &structures.ListNode{}


	for l1.Val != nil && l2.Val != nil {
		if l1 != nil {
			val1 := l1.Val
		}

		if l2.Val != nil {
			val2 := l2.Val
		}
	}


	return result
}
