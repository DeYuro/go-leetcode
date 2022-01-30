package problem23

import (
	"github.com/deyuro/go-leetcode/structures"
	"sort"
)

func mergeKLists(lists []*structures.ListNode) *structures.ListNode {

	if len(lists) < 1 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}
	var sorted []int

	for _,v := range lists {
		for v != nil {
			sorted = append(sorted, v.Val)
			v = v.Next
		}
	}

	sort.Ints(sorted)

	answer := &structures.ListNode{}
	head := answer

	for _,v := range sorted{
		answer.Next = &structures.ListNode{
			Val: v,
			Next: nil,
		}

		answer = answer.Next
	}

	return head.Next
}