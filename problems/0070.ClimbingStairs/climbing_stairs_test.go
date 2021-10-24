package problem70

import (
	"reflect"
	"testing"
)

type test struct {
	input int
	expected int
}

func getTestcases() []test {
	return []test{
		{
			2,
			2,
		},
		{
			3,
			3,
		},
		{
			5,
			8,
		},
	}
}

func TestClimbStairs(t *testing.T) {
	for _, v := range getTestcases() {
		output := climbStairs(v.input)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}