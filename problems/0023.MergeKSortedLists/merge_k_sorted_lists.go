package problem23

import (
	"container/heap"
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

type IntHeap []int

func mergeKLists2(lists []*structures.ListNode) *structures.ListNode {

	var h IntHeap

	for _, v := range lists {
		for v != nil {
			heap.Push(&h,v.Val)
			v = v.Next
		}
	}

	a := &structures.ListNode{}
	head := a
	for len(h) > 0 {
		a.Next = &structures.ListNode{
			Val: heap.Pop(&h).(int),
		}
		a = a.Next
	}

	return head.Next
}


func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(*h)
	x := old[n-1]
	*h = old[0:n-1]
	return x
}