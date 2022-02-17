package problem234

import "github.com/deyuro/go-leetcode/structures"

func isPalindrome(head *structures.ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	var s []int

	for head != nil {
		s = append(s,head.Val)
		head = head.Next
	}

	for i,j :=0,len(s)-1; i <= j; i,j = i+1,j-1 {
		if s[i] != s[j] {
			return false
		}
	}

	return true
}


func isPalindrome2(head *structures.ListNode) bool {
	length := Length(head)
	mid := (length - 1)/2

	// compute middle node
	midNode := head
	for count := 0; count < mid; count++ {
		midNode = midNode.Next
	}

	midHead := midNode.Next
	midNode.Next = nil

	reverseHead := reverseList(midHead)

	return compareLists(head, reverseHead)
}

func reverseList(head *structures.ListNode) *structures.ListNode {
	var reverseHead *structures.ListNode = nil

	for head != nil {
		temp := head.Next

		head.Next = reverseHead
		reverseHead = head

		head = temp
	}

	return reverseHead
}

func compareLists(head1, head2 *structures.ListNode) bool {
	for head1 != nil && head2 != nil {
		if head1.Val != head2.Val {
			return false
		}

		head1 = head1.Next
		head2 = head2.Next
	}

	return true
}

func Length(head *structures.ListNode) int {
	length := 0
	for ;head != nil; head, length = head.Next, length+1 {
	}

	return length
}