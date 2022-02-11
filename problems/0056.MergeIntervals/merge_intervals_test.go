package problem56

import (
	"reflect"
	"testing"
)

type test struct {
	input [][]int
	expected [][]int
}

func getTestcases() []test {
	return []test{
		{
			[][]int{{11,16},{1,3},{2,6},{8,10},{15,18}},
			[][]int{{1,6},{8,10},{11,18}},
		},
		{
			[][]int{{1,4},{2,3}},
			[][]int{{1,4}},
		},
		{
			[][]int{{1,3},{2,6},{8,10},{15,18}},
			[][]int{{1,6},{8,10},{15,18}},
		},
		{
			[][]int{{1,4},{4,5}},
			[][]int{{1,5}},
		},
	}
}


func TestMerge(t *testing.T)  {
	for _,v := range getTestcases() {
		output := merge(v.input)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}
func BenchmarkMerge(b *testing.B) {
	test := getTestcases()[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		merge(test.input)
	}
}
