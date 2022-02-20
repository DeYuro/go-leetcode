package temperature

import (
	"reflect"
	"testing"
)

type test struct {
	n []int
	expected []int
}

func getTestcases() []test {
	return []test{
		{
			[]int{13,12,15,11,9,12,16},
			[]int{2,1,4,2,1,1,0},
		},
	}
}

func TestTemperatures (t *testing.T) {
	for _,v := range getTestcases() {
		output := temperatures2(v.n)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}