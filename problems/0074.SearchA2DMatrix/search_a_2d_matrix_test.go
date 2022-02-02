package problem74

import (
	"reflect"
	"testing"
)

type test struct {
	input [][]int
	target int
	expected bool
}

func getTestcases() []test {
	return []test{
		{
			[][]int{{1,3,5,7},{10,11,16,20},{23,30,34,60}},
			3,
			true,
		},
		{
			[][]int{{1,3,5,7},{10,11,16,20},{23,30,34,60}},
			13,
			false,
		},
		{
			[][]int{{1}},
			1,
			true,
		},
	}
}

func TestSearchMatrix(t *testing.T) {
	for _, v := range getTestcases() {
		output := searchMatrix(v.input, v.target)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}

func BenchmarkSearchMatrix(b *testing.B) {
	test := getTestcases()[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		searchMatrix(test.input, test.target)
	}
}

func TestSearchMatrix2(t *testing.T) {
	for _, v := range getTestcases() {
		output := searchMatrix2(v.input, v.target)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}

func BenchmarkSearchMatrix2(b *testing.B) {
	test := getTestcases()[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		searchMatrix2(test.input, test.target)
	}
}