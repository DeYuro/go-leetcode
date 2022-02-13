package problem347

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
			[]int{1,1,1,2,2,3},
			2,
			[]int{1,2},
		},
		{
			[]int{1},
			1,
			[]int{1},
		},
	}
}
func TestTopKFrequent(t *testing.T) {
	for _, v := range getTestcases() {
		output := topKFrequent(v.input, v.k)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}

func BenchmarkTopKFrequent(b *testing.B) {
	test := getTestcases()[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		topKFrequent(test.input, test.k)
	}
}