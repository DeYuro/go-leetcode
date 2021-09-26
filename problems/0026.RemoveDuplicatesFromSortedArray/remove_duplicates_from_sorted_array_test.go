package problem26

import (
	"reflect"
	"testing"
)

type test struct {
	input    []int
	expected int
}

func getTestcases() []test {
	return []test{
		{
		[]int{1,1,2},
			2,
		},
		{
			[]int{0,0,1,1,1,2,2,3,3,4},
			5,
		},
	}
}

func TestRemoveDuplicates(t *testing.T) {
	for _, v := range getTestcases() {
		output := removeDuplicates(v.input)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}

func BenchmarkIsValid(b *testing.B) {
	test := getTestcases()[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		removeDuplicates(test.input)
	}
}