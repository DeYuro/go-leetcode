package problem206

import (
	"github.com/deyuro/go-leetcode/structures"
	"reflect"
	"testing"
)

type test struct {
	head, expected *structures.ListNode
}

func getTestcases() []test {
	return []test{
		{
			structures.CreateListNode([]int{1,2,3,4,5}),
			structures.CreateListNode([]int{5,4,3,2,1}),
		},
		{
			structures.CreateListNode([]int{1,2}),
			structures.CreateListNode([]int{2,1}),
		},
		{
			structures.CreateListNode([]int{}),
			structures.CreateListNode([]int{}),
		},
	}
}

func TestReverseList(t *testing.T)  {
	for _,v := range getTestcases() {
		output := reverseList(v.head)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}

func BenchmarkMergeTwoLists(b *testing.B)  {
	test := getTestcases()[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverseList(test.head)
	}
}