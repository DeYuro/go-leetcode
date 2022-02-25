package problem92

import (
	"github.com/deyuro/go-leetcode/structures"
	"reflect"
	"testing"
)

type test struct {
	head, expected *structures.ListNode
	left,right int
}

func getTestcases() []test {
	return []test{
		{
			structures.CreateListNode([]int{1,2,3,4,5}),
			structures.CreateListNode([]int{1,4,3,2,5}),
			2,
			4,
		},
		{
			structures.CreateListNode([]int{5}),
			structures.CreateListNode([]int{5}),
			1,
			1,
		},

	}
}

func TestReverseList(t *testing.T)  {
	for _,v := range getTestcases() {
		output := reverseBetween(v.head,v.left,v.right)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}

func BenchmarkMergeTwoLists(b *testing.B)  {
	test := getTestcases()[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverseBetween(test.head, test.left, test.right)
	}
}