package problem7

import (
	"reflect"
	"testing"
)

type test struct {
	input    int
	expected int
}

func TestReverseInteger(t *testing.T)  {
	for _,v := range getTestcases() {
		output := reverse(v.input)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}

func BenchmarkReverseInteger(b *testing.B)  {
	test := getTestcases()[2]
	b.ResetTimer()
	for i:=0; i < b.N; i++ {
		reverse(test.input)
	}
}

func getTestcases() []test {
	return []test{
		{
			123,
			321,
		},
		{
			-123,
			-321,
		},
		{
			120,
			21,
		},
		{
			0,
			0,
		},
		{
			1534236469,
			0,
		},
		{
			-2147483648,
			0,
		},
	}
}