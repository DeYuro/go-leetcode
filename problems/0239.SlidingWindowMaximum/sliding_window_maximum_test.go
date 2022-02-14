package problem239

import (
	"reflect"
	"testing"
)

type test struct {
	input    []int
	k int
	expected []int
}

func getTestcases() []test {
	return []test{
		{
			[]int{8,7,6,9},
			2,
			[]int{8,7,9},
		},
		{
			[]int{1,3,-1,-3,5,3,6,7},
			3,
			[]int{3,3,5,5,6,7},
		},
		{
			[]int{1},
			1,
			[]int{1},
		},
	}
}
func TestMaxSlidingWindow(t *testing.T) {
	for _, v := range getTestcases() {
		output := maxSlidingWindow(v.input, v.k)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}

func BenchmarkMaxSlidingWindow(b *testing.B) {
	test := getTestcases()[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maxSlidingWindow(test.input, test.k)
	}
}