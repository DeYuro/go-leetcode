package problem480

import (
	"reflect"
	"testing"
)

type test struct {
	input    []int
	k int
	expected []float64
}

func getTestcases() []test {
	return []test{
		{
			[]int{1,3,-1,-3,5,3,6,7},
			3,
			[]float64{1.00000,-1.00000,-1.00000,3.00000,5.00000,6.00000},
		},
		{
			[]int{1,2,3,4,2,3,1,4,2},
			3,
			[]float64{2.00000,3.00000,3.00000,3.00000,2.00000,3.00000,2.00000},
		},
	}
}
func TestMaxSlidingWindow(t *testing.T) {
	for _, v := range getTestcases() {
		output := medianSlidingWindow(v.input, v.k)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}

func BenchmarkMaxSlidingWindow(b *testing.B) {
	test := getTestcases()[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		medianSlidingWindow(test.input, test.k)
	}
}