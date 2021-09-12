package problem21

import (
	"github.com/go-leetcode/structures"
	"sort"
)

func mergeTwoLists(l1 *structures.ListNode,l2 *structures.ListNode) *structures.ListNode {
	sorted := []int{}
	merged := &structures.ListNode{}
	head := merged
	for l1 != nil {
		sorted = append(sorted, l1.Val)
		l1 = l1.Next
	}

	for l2 != nil {
		sorted = append(sorted, l2.Val)
		l2 = l2.Next
	}

	sort.Ints(sorted)

	for _,v := range sorted{
		merged.Next = &structures.ListNode{
			Val:  v,
			Next: nil,
		}
		merged = merged.Next
	}

	return head.Next
}