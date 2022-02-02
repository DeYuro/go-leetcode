package problem704

import (
	"reflect"
	"testing"
)

type test struct {
	input []int
	target int
	expected int
}

func getTestcases() []test {
	return []test{
		{
			[]int{-1,0,3,5,9,12},
			9,
			4,
		},
		{
			[]int{-1,0,3,5,9,12},
			2,
			-1,
		},
	}
}

func TestBinarySearch(t *testing.T) {
	for _, v := range getTestcases() {
		output := search(v.input, v.target)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	test := getTestcases()[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		search(test.input, test.target)
	}
}
