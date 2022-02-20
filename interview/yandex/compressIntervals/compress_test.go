package compress

import (
	"reflect"
	"testing"
)

type test struct {
	input    []int
	expected string
}

func getTestcases() []test {
	return []test{
		{
			[]int{4,1,3,2,6,8,7,11,0},
			"0-4,6-8,11",
		},
		{
			[]int{4,1,3,2,6,8,7,11,12},
			"1-4,6-8,11-12",
		},
		{
			[]int{1,4},
			"1,4",
		},
		{
			[]int{1},
			"1",
		},
		{
			[]int{},
			"",
		},
	}
}

func TestCompress(t *testing.T) {
	for _, v := range getTestcases() {
		output := compress(v.input)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}
