package problem154

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
			[]int{1,3,5},
			1,
		},
		{
			[]int{2,2,2,0,1},
			0,
		},
		{
			[]int{1,0,1,1,1},
			0,
		},
		{
			[]int{1,3},
			1,
		},
	}
}


func TestFindMin(t *testing.T) {
	for _, v := range getTestcases() {
		output := findMin(v.input)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}

func BenchmarkFindMin(b *testing.B) {
	test := getTestcases()[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findMin(test.input)
	}
}
