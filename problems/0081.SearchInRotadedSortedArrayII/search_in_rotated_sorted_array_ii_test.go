package problem81

import (
	"reflect"
	"testing"
)

type test struct {
	input    []int
	target   int
	expected bool
}

func getTestcases() []test {
	return []test{
		{
			[]int{1,0,1,1,1},
			0,
			true,
		},
		{
			[]int{4,5,6,6,7,7,0,1,2,4,4},
			8,
			false,
		},
		{
			[]int{4,5,6,6,7,0,1,2,4,4},
			8,
			false,
		},
		{
			[]int{4,5,6,6,7,0,1,2,4,4},
			0,
			true,
		},
		{
			[]int{4,5,6,6,7,0,1,2,4,4},
			3,
			false,
		},
		{
			[]int{2,5,6,0,0,1,2},
			0,
			true,
		},
		{
			[]int{2,5,6,0,0,1,2},
			3,
			false,
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
