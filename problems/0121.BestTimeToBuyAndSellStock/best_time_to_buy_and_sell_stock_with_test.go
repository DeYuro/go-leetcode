package problem121

import (
	"reflect"
	"testing"
)

type test struct {
	input []int
	expected int
}

func getTestcases() []test {
	return []test{
		{
			[]int{7,1,5,3,6,4},
			5,
		},
		{
			[]int{7,6,4,3,1},
			0,
		},
	}
}
func TestMaxProfit(t *testing.T) {
	for _, v := range getTestcases() {
		output := maxProfit(v.input)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}

func BenchmarkMaxProfit(b *testing.B) {
	test := getTestcases()[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maxProfit(test.input)
	}
}