package problem200

import (
	"reflect"
	"testing"
)



type test struct {
	input    [][]byte
	expected int
}

func getTestcases() []test {
	return []test{
		{
			[][]byte{
				{'1', '0', '1', '1', '1'},
				{'1', '0', '1', '0', '1'},
				{'1', '1', '1', '0', '1'}},
			1,
		},
		{
			[][]byte{
				{'1','1','1','1','0'},
				{'1','1','0','1','0'},
				{'1','1','0','0','0'},
				{'0','0','0','0','0'}},
			1,
		},
		{
			[][]byte{
				{'1','1','0','0','0'},
				{'1','1','0','0','0'},
				{'0','0','1','0','0'},
				{'0','0','0','1','1'}},
			3,
		},
	}
}

func TestNumIslands(t *testing.T) {
	for _, v := range getTestcases() {
		output := numIslands(v.input)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}

func BenchmarkNumIslands(b *testing.B) {
	test := getTestcases()[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		numIslands(test.input)
	}
}