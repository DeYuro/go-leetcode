package problem67

import (
	"reflect"
	"testing"
)

type test struct {
	a string
	b string
	expected string
}

func getTestcases() []test {
	return []test{
		{
			"11",
			"1",
			"100",
		},
		{
			"1010",
			"1011",
			"10101",
		},
	}
}

func TestAddBinary(t *testing.T) {
	for _, v := range getTestcases() {
		output := addBinary(v.a, v.b)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}

func BenchmarkAddBinary(b *testing.B) {
	test := getTestcases()[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		addBinary(test.a, test.b)
	}
}