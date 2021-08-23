package problem2

import (
	"github.com/go-leetcode/structures"
)

func addTwoNumbers(l1, l2 *structures.ListNode)  *structures.ListNode {
	var mem int
	var list []int
	for {
		if l1 == nil && l2 == nil {
			break
		}

		var num1 int
		var num2 int
		if l1 != nil {
			num1 = l1.Val
		}

		if l2 != nil {
			num2 = l2.Val
		}

		result := num1 + num2 + mem

		if result > 9 {
			mem = 1
			result = result % 10
		} else {
			mem = 0
		}

		list = append(list, result)

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	if mem == 1 {
		list = append(list, 1)
	}

	return structures.CreateListNode(list)
}