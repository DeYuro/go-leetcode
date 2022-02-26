package problem989

import (
	"reflect"
	"testing"
)

type test struct {
	num []int
	k     int
	expected []int
}

func getTestcases() []test {
	return []test{
		{
			[]int{0},
			10000,
			[]int{1,0,0,0,0},
		},
		{
			[]int{0},
			23,
			[]int{2,3},
		},
		{
			[]int{1,2,0,0},
			34,
			[]int{1,2,3,4},
		},
		{
			[]int{2,7,4},
			181,
			[]int{4,5,5},
		},
		{
			[]int{2,1,5},
			806,
			[]int{1,0,2,1},
		},
	}
}

func TestAddToArrayForm(t *testing.T) {
	for _, v := range getTestcases() {
		output := addToArrayForm(v.num, v.k)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}

func TestAddToArrayForm2(t *testing.T) {
	for _, v := range getTestcases() {
		output := addToArrayForm2(v.num, v.k)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}

func BenchmarkAddToArrayForm(b *testing.B)  {
	test := getTestcases()[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		addToArrayForm(test.num, test.k)
	}
}

func BenchmarkAddToArrayForm2(b *testing.B)  {
	test := getTestcases()[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		addToArrayForm2(test.num, test.k)
	}
}