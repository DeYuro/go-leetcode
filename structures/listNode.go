package structures

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateListNode(input []int) *ListNode {
	var list *ListNode
	var cursor *ListNode
	var next *ListNode
	for _,v := range input {
		if list == nil {
			list = &ListNode{
				Val:  v,
				Next: nil,
			}

			cursor = list
			continue
		}

		next = &ListNode{
			v,
			nil,
		}

		cursor.Next = next
		cursor = next
	}

	return list
}

func (b ListNode) String() string {
	str := ""
	p := &b
	for p != nil {
		str +=  strconv.Itoa(p.Val)
		p = p.Next
	}
	return fmt.Sprintf("%s", str)
}