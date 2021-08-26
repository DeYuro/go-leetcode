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
	input []int
	target int
	expected []int
}

func getTestcases() []test {
	return []test{
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