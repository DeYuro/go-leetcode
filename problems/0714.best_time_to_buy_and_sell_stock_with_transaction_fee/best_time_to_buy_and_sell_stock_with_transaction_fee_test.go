package problem714

import (
	"reflect"
	"testing"
)

type test struct {
	input    []int
	fee 	 int
	expected int
}

func getTestcases() []test {
	return []test{
		{
			[]int{1,3,2,8,4,9},
			2,
			8,
		},
		{
			[]int{1,3,7,5,10,3},
			3,
			6,
		},
	}
}
func TestMaxProfit(t *testing.T) {
	for _, v := range getTestcases() {
		output := maxProfit(v.input,v.fee)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}

func BenchmarkMaxProfit(b *testing.B) {
	test := getTestcases()[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maxProfit(test.input,test.fee)
	}
}

