package problem21

import (
	"github.com/go-leetcode/structures"
	"reflect"
	testing "testing"
)

type test struct {
	l1,l2,expected *structures.ListNode
}

func getTestcases() []test {
	return []test{
		{
			structures.CreateListNode([]int{1}),
			structures.CreateListNode([]int{0,3,4}),
			structures.CreateListNode([]int{0,1,3,4}),
		},
		{
			structures.CreateListNode([]int{1,2,4}),
			structures.CreateListNode([]int{1,3,4}),
			structures.CreateListNode([]int{1,1,2,3,4,4}),
		},
		{
			structures.CreateListNode([]int{}),
			structures.CreateListNode([]int{}),
			structures.CreateListNode([]int{}),
		}, {
			structures.CreateListNode([]int{}),
			structures.CreateListNode([]int{0}),
			structures.CreateListNode([]int{0}),
		},
	}
}

func TestMergeTwoLists(t *testing.T)  {
	for _,v := range getTestcases() {
		output := mergeTwoLists(v.l1, v.l2)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}