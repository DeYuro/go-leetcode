package problem53

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
			[]int{-2,1,-3,4,-1,2,1,-5,4},
			6,
		},
		{
			[]int{5,4,-1,7,8},
			23,
		},
		{
			[]int{1},
			1,
		},
		{
			[]int{-2, 1},
			1,
		},
		{
			[]int{-2, -1},
			-1,
		},
		{

		[]int{-1,0,-2},
		0,
		},
	}

}

func TestMaxSubArray(t *testing.T) {
	for _, v := range getTestcases() {
		output := maxSubArray(v.input)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}

func BenchmarkMaxSubArray(b *testing.B) {
	test := getTestcases()[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maxSubArray(test.input)
	}
}