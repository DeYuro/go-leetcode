package problem92

import "github.com/deyuro/go-leetcode/structures"

func reverseBetween(head *structures.ListNode, left int, right int) *structures.ListNode {
	res := &structures.ListNode{}
	res.Next = head
	head = res
	for i := 0; i < left-1 ; i++ {
		head = head.Next
	}
	var prev, next *structures.ListNode
	cur := head.Next
	for i := 0; i <= right-left; i++ {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	head.Next = prev
	for head != nil {
		if head.Next == nil {
			head.Next = cur
			break
		}
		head = head.Next
	}
	return res.Next
}