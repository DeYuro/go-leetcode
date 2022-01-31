package problem206

import "github.com/deyuro/go-leetcode/structures"

func reverseList(head *structures.ListNode) *structures.ListNode {
	var cur, next, prev *structures.ListNode
	cur = head

	for cur != nil {
		//[1->2->3->4->5]

		// [2->3->4->5->nil]
		// [3->4->5->nil]
		// [4->5->nil]
		// [5->nil]
		// [nil]
		next = cur.Next

		// [nil] on first
		// [1->nil]
		// [2->1->nil]
		// [3->2->1->nil]
		// [4->3->2->1->nil]
		cur.Next = prev

		// [1->nil]
		// [2->1->nil]
		// [3->2->1->nil]
		// [4->3->2->1->nil]
		// [5->4->3->2->1]  !!! prev is answer
		prev = cur

		// [2->3->4->5->nil]
		// [3->4->5->nil]
		// [4->5->nil]
		// [5->nil]
		// [nil] loop end
		cur  = next
	}

	return prev
}

