package problem1

import (
	"reflect"
	"testing"
)

type test struct {
	input []int
	target int
	expected []int
}

func getTestcases() []test {
	return []test{
		{
			[]int{2,7,11,15},
				9,
				[]int{0,1},
		},
		{
			[]int{3,2,4},
				6,
				[]int{1,2},
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
	test := getTestcases()[0]
	b.ResetTimer()
	for i:=0; i < b.N; i++ {
		twoSum(test.input,test.target)
	}
}

func BenchmarkTwoSumLoop(b *testing.B) {
	test := getTestcases()[0]
	b.ResetTimer()
	for i:=0; i < b.N; i++ {
		twoSumLoops(test.input,test.target)
	}
}