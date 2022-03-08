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
			[]int{2,2,3,2},
			3,
		},
		{
			[]int{0,1,0,1,0,1,99},
			99,
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
func BenchmarkSingleNumber(b *testing.B) {
	test := getTestcases()[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		singleNumber(test.input)
	}
}

func BenchmarkSingleNumber2(b *testing.B) {
	test := getTestcases()[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		singleNumber2(test.input)
	}
}