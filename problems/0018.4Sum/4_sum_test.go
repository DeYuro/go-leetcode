package problem14

import (
	"testing"
)

type test struct {
	input []int
	target int
	expected [][]int
}

func getTestcases() []test {
	return []test{
		{
			[]int{1,0,-1,0,-2,2},
			0,
			[][]int{{-2,-1,1,2},{-2,0,0,2},{-1,0,0,1}},
		},
		{
			[]int{2,2,2,2},
			8,
			[][]int{{2,2,2,2}},
		},
	}
}

func TestFourSum(t *testing.T)  {
	for _,v := range getTestcases() {
		output := fourSum(v.input, v.target)
		if len(output) != len(v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}