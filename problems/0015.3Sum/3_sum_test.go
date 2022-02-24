package problem15

import (
	"testing"
)

type test struct {
	input []int
	expected [][]int
}

func getTestcases() []test {
	return []test{
		{
			[]int{-1,0,1,2,-1,-4},
			[][]int{{-1,-1,2},{-1,0,1}},
		},
		{
			[]int{},
			[][]int{},
		},
		{
			[]int{0},
			[][]int{},
		},
	}
}

func TestThreeSum(t *testing.T)  {
	for _,v := range getTestcases() {
		output := threeSum(v.input)
		if len(output) != len(v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}

func BenchmarkThreeSum(b *testing.B)  {
	test := getTestcases()[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		threeSum(test.input)
	}
}