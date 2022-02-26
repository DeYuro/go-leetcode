package problem2095

import "github.com/deyuro/go-leetcode/structures"

func deleteMiddle(head *structures.ListNode) *structures.ListNode {
	if head.Next == nil {
		return nil
	}
	tmp := head
	c := -1
	for tmp != nil {
		c++
		tmp = tmp.Next
	}

	middle := c / 2
	if c % 2 != 0 {
		middle++
	}
	c = 0
	res := head
	for head != nil {
		if c+1 == middle {
			head.Next = head.Next.Next
		}

		c++
		head = head.Next
	}

	return res
}
