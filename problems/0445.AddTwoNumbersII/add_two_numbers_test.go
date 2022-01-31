package problem445

import (
	"github.com/deyuro/go-leetcode/structures"
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
			structures.CreateListNode([]int{7,2,4,3}),
			structures.CreateListNode([]int{5,6,4}),
			structures.CreateListNode([]int{7,8,0,7}),
		},
		{
			structures.CreateListNode([]int{0}),
			structures.CreateListNode([]int{0}),
			structures.CreateListNode([]int{0}),
		},
		{
			structures.CreateListNode([]int{2,4,3}),
			structures.CreateListNode([]int{5,6,4}),
			structures.CreateListNode([]int{8,0,7}),
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

func BenchmarkAddTwoNumbers(b *testing.B) {
	test := getTestcases()[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		addTwoNumbers(test.l1,test.l2)
	}
}

