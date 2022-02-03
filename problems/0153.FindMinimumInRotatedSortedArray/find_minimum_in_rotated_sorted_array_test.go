package problem153

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
			[]int{2,4,5,6,7,0,1},
			0,
		},
		{
			[]int{3,4,5,1,2},
			1,
		},
		{
			[]int{11,13,15,17},
			11,
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
func TestFindMin2(t *testing.T) {
	for _, v := range getTestcases() {
		output := findMin2(v.input)
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
