package problem445

import (
	"github.com/deyuro/go-leetcode/structures"
)


func addTwoNumbers(l1,l2 *structures.ListNode) *structures.ListNode {
	listNode := &structures.ListNode{}
	mem := 0
	head := listNode

	l1 = reverse(l1)
	l2 = reverse(l2)

	for l1 != nil || l2 != nil || mem != 0 {
		res := 0
		if l1 != nil {
			res+=l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			res+=l2.Val
			l2 = l2.Next
		}
		res += mem
		mem = res / 10

		listNode.Next = &structures.ListNode{
			Val: res % 10,
		}
		listNode = listNode.Next
	}

	return reverse(head.Next)
}

func reverse(head *structures.ListNode) *structures.ListNode{

	var cur,next,prev *structures.ListNode

	cur = head
	for cur != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	return prev
}