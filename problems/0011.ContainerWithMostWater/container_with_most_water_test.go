package problem11

import (
	"reflect"
	"testing"
)

type test struct {
	input    []int
	expected int
}

func getTestcases() []test {
	return []test{
		{
			[]int{1,8,6,2,5,4,8,3,7},
			49,
		},
		{
			[]int{4,3,2,1,4},
			16,
		},
		{
			[]int{1,2,1},
			2,
		},
		{

		[]int{1,2},
		1,
		},
	}
}

func TestLongestCommonPrefix(t *testing.T) {
	for _, v := range getTestcases() {
		output := maxArea(v.input)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}