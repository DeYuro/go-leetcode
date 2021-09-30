package problem35

import (
	"reflect"
	"testing"
)

type test struct {
	input    []int
	target   int
	expected int
}

func getTestcases() []test {
	return []test{
		{
		[]int{2,3,5,6,9},
		7,
		4,
		},
		{
			[]int{1,3,5,6},
			5,
			2,
		},
		{
			[]int{1,3,5,6},
			2,
			1,
		},
		{
			[]int{1,3,5,6},
			7,
			4,
		},
		{
			[]int{1,3,5,6},
			0,
			0,
		},
		{
			[]int{1},
			0,
			0,
		},
		{
			[]int{1,2,5},
			4,
			2,
		},
		{
			[]int{1,3},
			2,
			1,
		},
	}
}

func TestSearchInsert(t *testing.T) {
	for _, v := range getTestcases() {
		output := searchInsert(v.input, v.target)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}

func TestSearchInsert2(t *testing.T) {
	for _, v := range getTestcases() {
		output := searchInsert2(v.input, v.target)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}

func TestSearchInsert3(t *testing.T) {
	for _, v := range getTestcases() {
		output := searchInsert3(v.input, v.target)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}

func BenchmarkSearchInsert(b *testing.B) {
	test := getTestcases()[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		searchInsert(test.input, test.target)
	}
}

func BenchmarkSearchInsert2(b *testing.B) {
	test := getTestcases()[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		searchInsert2(test.input, test.target)
	}
}

func BenchmarkSearchInsert3(b *testing.B) {
	test := getTestcases()[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		searchInsert3(test.input, test.target)
	}
}