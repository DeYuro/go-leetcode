package problem6

import (
	"reflect"
	"testing"
)

type test struct {
	input    string
	numRows  int
	expected string
}

func getTestcases() []test {
	return []test{
		{
			`PAYPALISHIRING`,
			3,
			`PAHNAPLSIIGYIR`,
		},
		{
			"ABC",
			1,
			"HZ",
		},
	}
}

func TestZigZagConvert(t *testing.T)  {
	for _,v := range getTestcases() {
		output := convert(v.input, v.numRows)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}

func TestZigZagConvert2(t *testing.T)  {
	for _,v := range getTestcases() {
		output := convert2(v.input, v.numRows)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}