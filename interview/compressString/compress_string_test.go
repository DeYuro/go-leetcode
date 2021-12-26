package compressString

import (
	"reflect"
	"testing"
)

type test struct {
	input    string
	expected string
}

func getTestcases() []test {
	return []test{
		{
			"aabbccc",
			"a2b2c3",
		},
		{
			"a",
			"a",
		},
		{
			"abbbbbbbbbbbb",
			"ab12",
		},
	}
}

func TestCompress(t *testing.T) {
	for _, v := range getTestcases() {
		output := compressString(v.input)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}
