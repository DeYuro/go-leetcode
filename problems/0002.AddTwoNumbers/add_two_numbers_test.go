package problem2

import (
	"github.com/go-leetcode/structures"
	"reflect"
	"testing"
)

type test struct {
	l1       *structures.ListNode
	l2       *structures.ListNode
	expected *structures.ListNode
}

func getTestcases() []test {
	return []test{
		{
			structures.CreateListNode([]int{2,4,3}),
			structures.CreateListNode([]int{5,6,4}),
			structures.CreateListNode([]int{7,0,8}),
		},
		{
			structures.CreateListNode([]int{0}),
			structures.CreateListNode([]int{0}),
			structures.CreateListNode([]int{0}),
		},
		{
			structures.CreateListNode([]int{9,9,9,9,9,9,9}),
			structures.CreateListNode([]int{9,9,9,9}),
			structures.CreateListNode([]int{8,9,9,9,0,0,0,1}),
		},{
			structures.CreateListNode([]int{1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1}),
			structures.CreateListNode([]int{5,6,4}),
			structures.CreateListNode([]int{6,6,4,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1}),
		},
	}
}

func TestAddTwoNumbers(t *testing.T)  {
	for _,v := range getTestcases() {
		output := addTwoNumbers(v.l1, v.l2)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}