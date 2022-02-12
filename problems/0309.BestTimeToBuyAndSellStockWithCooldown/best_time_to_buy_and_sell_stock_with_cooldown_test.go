package problem309

import (
	"reflect"
	"testing"
)

type test struct {
	input []int
	expected int
}

func getTestcases() []test {
	return []test{
		{
			[]int{1,2,2,2,4,5},
			4,
		},
		{
			[]int{1,2,3,0,2},
			3,
		},
		{
			[]int{1},
			0,
		},
		{
			[]int{1,2,4},
			3,
		},
	}
}
func TestMaxProfit(t *testing.T) {
	for _, v := range getTestcases() {
		output := maxProfit(v.input)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}

func TestMaxProfit2(t *testing.T) {
	for _, v := range getTestcases() {
		output := maxProfit2(v.input)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}


func TestMaxProfit3(t *testing.T) {
	for _, v := range getTestcases() {
		output := maxProfit3(v.input)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}
