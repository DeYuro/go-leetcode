package problem136

import (
	"reflect"
	"testing"
)

type test struct {
	input     []int
	expected int
}

func getTestcases() []test {
	return []test{
		{
			[]int{2,2,1},
			1,
		},
		{
			[]int{4,1,2,1,2},
			4,
		},
		{
			[]int{1},
			1,
		},

	}
}

func TestSingleNumber(t *testing.T) {
	for _, v := range getTestcases() {
		output := singleNumber(v.input)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}

func TestSingleNumber2(t *testing.T) {
	for _, v := range getTestcases() {
		output := singleNumber2(v.input)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}