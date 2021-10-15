package problem66

import (
	"reflect"
	"testing"
)

type test struct {
	input    []int
	expected []int
}

func TestIsValid(t *testing.T) {
	for _, v := range getTestcases() {
		output := plusOne(v.input)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}

func getTestcases() []test {
	return []test{
		{
			[]int{1,2,3},
			[]int{1,2,4},
		},
		{
			[]int{1,2,3},
			[]int{1,2,4},
		}, {
			[]int{4,3,2,1},
			[]int{4,3,2,2},
		}, {
			[]int{0},
			[]int{1},
		},
		{
			[]int{9},
			[]int{1,0},
		},
	}
}

func BenchmarkPlusOne(b *testing.B) {
	test := getTestcases()[2]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		plusOne(test.input)
	}
}