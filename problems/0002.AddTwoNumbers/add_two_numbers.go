package problem2

import (
	"github.com/deyuro/go-leetcode/structures"
)


func addTwoNumbers(l1,l2 *structures.ListNode) *structures.ListNode {
	var mem int
	listNode := &structures.ListNode{}
	var head = listNode
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

		head.Next = &structures.ListNode{
			Val: res % 10,
		}
		head = head.Next
		mem = res / 10
	}

	return listNode.Next
}

func addTwoNumbers2(l1,l2 *structures.ListNode) *structures.ListNode {
	return recursion(l1,l2,0)
}

func recursion(l1, l2 *structures.ListNode, sum int) *structures.ListNode {
	if l1 == nil && l2 == nil && sum == 0 {
		return nil
	}
	if l1 != nil {
		sum += l1.Val
		l1 = l1.Next
	}

	if l2 != nil {
		sum += l2.Val
		l2 = l2.Next
	}

	return &structures.ListNode{
		Val: sum % 10,
		Next: recursion(l1, l2, sum/10),
	}
}

func addTwoNumbers3(l1, l2 *structures.ListNode)  *structures.ListNode {
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