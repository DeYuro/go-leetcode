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

func TestClimbStairsDp(t *testing.T) {
	for _, v := range getTestcases() {
		output := climbStairsDp(v.input)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}

func TestClimbStairsFib(t *testing.T) {
	for _, v := range getTestcases() {
		output := climbStairsFib(v.input)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}

func BenchmarkClimbStairsDp(b *testing.B) {
	test := getTestcases()[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		climbStairsDp(test.input)
	}
}

func BenchmarkClimbStairsFib(b *testing.B) {
	test := getTestcases()[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		climbStairsFib(test.input)
	}
}