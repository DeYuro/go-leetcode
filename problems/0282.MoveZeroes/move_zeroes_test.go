package problem283

import (
	"reflect"
	"testing"
)

type test struct {
	input    []int
	expected []int // not in leetcode
}

func getTestcases() []test {
	return []test{
		{
			[]int{0,1,0,3,12},
			[]int{1,3,12,0,0},
		},
		{
			[]int{0,0,1},
			[]int{1,0,0},
		},
		{
			[]int{0},
			[]int{0},
		},
	}
}

func TestMoveZeroes(t *testing.T) {
	for _, v := range getTestcases() {
		output := moveZeroes(v.input)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}

func TestMoveZeroes2(t *testing.T) {
	for _, v := range getTestcases() {
		output := moveZeroes2(v.input)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}

func TestMoveZeroes3(t *testing.T) {
	for _, v := range getTestcases() {
		output := moveZeroes3(v.input)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}

func BenchmarkMoveZeroes(b *testing.B) {
	test := getTestcases()[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		moveZeroes(test.input)
	}
}

func BenchmarkMoveZeroes2(b *testing.B) {
	test := getTestcases()[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		moveZeroes2(test.input)
	}
}

func BenchmarkMoveZeroes3(b *testing.B) {
	test := getTestcases()[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		moveZeroes3(test.input)
	}
}