package problem27

import (
	"reflect"
	"testing"
)

type test struct {
	nums     []int
	val      int
	expected int
}

func getTestcases() []test {
	return []test{
		{
			[]int{3, 2, 2, 3},
			3,
			2,
		},
		{
			[]int{0, 1, 2, 2, 3, 0, 4, 2},
			2,
			5,
		}}

}

func TestRemoveElement(t *testing.T) {
	for _, v := range getTestcases() {
		output := removeElement(v.nums, v.val)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}

func TestRemoveElement2(t *testing.T) {
	for _, v := range getTestcases() {
		output := removeElement(v.nums, v.val)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}

// TODO BENCH