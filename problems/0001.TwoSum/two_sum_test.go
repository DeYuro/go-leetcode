package problem1

import (
	"reflect"
	"testing"
)

 type ListNode struct {
	     Val int
	     Next *ListNode
 }

type test struct {
	l1 *ListNode
	l2 *ListNode
	expected *ListNode
}

func getTestcases() []test {
	return []test{
		{
			&ListNode{
				Val:  2,
				Next: &ListNode{
					Val:  4,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
			&ListNode{
				Val:  5,
				Next: &ListNode{
					Val:  6,
					Next: &ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
			&ListNode{
				Val:  7,
				Next: &ListNode{
					Val:  0,
					Next: &ListNode{
						Val:  8,
						Next: nil,
					},
				},
			},
		},
		{
			[]int{3,2,4},
				6,
				[]int{1,2},
		}, {
			[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 6, 7},
			13,
			[]int{50, 51},
		},
	}
}

func TestTwoSum(t *testing.T) {
	for _,v := range getTestcases() {
		output := twoSum(v.input,v.target)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}

func TestTwoSumLoop(t *testing.T) {
	for _,v := range getTestcases() {
		output := twoSumLoops(v.input,v.target)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}

func BenchmarkTwoSum(b *testing.B) {
	test := getTestcases()[2]
	b.ResetTimer()
	for i:=0; i < b.N; i++ {
		twoSum(test.input,test.target)
	}
}

func BenchmarkTwoSumLoop(b *testing.B) {
	test := getTestcases()[2]
	b.ResetTimer()
	for i:=0; i < b.N; i++ {
		twoSumLoops(test.input,test.target)
	}
}