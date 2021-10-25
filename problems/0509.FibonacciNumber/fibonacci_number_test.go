package problem509

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
			1,
		},
		{
			3,
			2,
		},
		{
			4,
			3,
		},
	}
}

func TestFibonacciNumber(t *testing.T) {
	for _, v := range getTestcases() {
		output := fib(v.input)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}