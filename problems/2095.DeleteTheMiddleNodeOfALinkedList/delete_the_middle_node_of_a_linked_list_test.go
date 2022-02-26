package problem2095

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
			structures.CreateListNode([]int{1,3,4,7,1,2,6}),
			structures.CreateListNode([]int{1,3,4,1,2,6}),
		},
		{
			structures.CreateListNode([]int{1,2,3,4}),
			structures.CreateListNode([]int{1,2,4}),
		},
		{
			structures.CreateListNode([]int{2,1}),
			structures.CreateListNode([]int{2}),
		},
	}
}

func TestDeleteMiddle(t *testing.T)  {
	for _,v := range getTestcases() {
		output := deleteMiddle(v.head)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}

func BenchmarkDeleteMiddle(b *testing.B)  {
	test := getTestcases()[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		deleteMiddle(test.head)
	}
}