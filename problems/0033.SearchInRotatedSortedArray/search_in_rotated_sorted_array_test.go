package problem33

import (
	"reflect"
	"testing"
)

type test struct {
	input    []int
	target   int
	expected int
}

func getTestcases() []test {
	return []test{
		{
			[]int{4, 5, 6, 7, 8, 0, 1, 2},
			8,
			4,
		},
		{
			[]int{3,1},
			3,
			0,
		},
		{
			[]int{3,1},
			0,
			-1,
		},
		{
			[]int{4, 5, 6, 7, 0, 1, 2},
			0,
			4,
		},
		{
			[]int{4, 5, 6, 7, 0, 1, 2},
			3,
			-1,
		},
		{
			[]int{1},
			0,
			-1,
		},
	}
}

func TestSearch(t *testing.T) {
	for _, v := range getTestcases() {
		output := search(v.input, v.target)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}

func BenchmarkSearch(b *testing.B) {
	test := getTestcases()[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		search(test.input, test.target)
	}
}
