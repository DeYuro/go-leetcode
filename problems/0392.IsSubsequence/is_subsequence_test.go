package problem392

import (
	"reflect"
	"testing"
)

type test struct {
	input    string
	target   string
	expected bool
}

func getTestcases() []test {
	return []test{
		{
			"abc",
			"ahbgdc",
			true,
		},
		{
			"axc",
			"ahbgdc",
			false,
		},
	}
}

func TestIsSubsequence(t *testing.T) {
	for _, v := range getTestcases() {
		output := isSubsequence(v.input, v.target)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}

func BenchmarkIsSubsequence(b *testing.B) {
	test := getTestcases()[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isSubsequence(test.input, test.target)
	}
}