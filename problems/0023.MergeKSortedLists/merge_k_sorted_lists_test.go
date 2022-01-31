package problem23

import (
	"github.com/deyuro/go-leetcode/structures"
	"reflect"
	"testing"
)

type test struct {
	input []*structures.ListNode
	expected *structures.ListNode
}

func getTestcases() []test {
	return []test{
		{
			[]*structures.ListNode{},
			structures.CreateListNode([]int{}),
		},
		{
			[]*structures.ListNode{structures.CreateListNode([]int{})},
			structures.CreateListNode([]int{}),
		},
		{
			[]*structures.ListNode{
				structures.CreateListNode([]int{1,4,5}),
				structures.CreateListNode([]int{1,3,4}),
				structures.CreateListNode([]int{2,6}),
			},
			structures.CreateListNode([]int{1,1,2,3,4,4,5,6}),
		},
	}
}

func TestMergeKLists(t *testing.T)  {
	for _,v := range getTestcases() {
		output := mergeKLists(v.input)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}

func TestMergeKLists2(t *testing.T)  {
	for _,v := range getTestcases() {
		output := mergeKLists2(v.input)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}

func BenchmarkMergeKLists(b *testing.B)  {
	test := getTestcases()[2]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mergeKLists(test.input)
	}
}

func BenchmarkMergeKLists2(b *testing.B)  {
	test := getTestcases()[2]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mergeKLists2(test.input)
	}
}