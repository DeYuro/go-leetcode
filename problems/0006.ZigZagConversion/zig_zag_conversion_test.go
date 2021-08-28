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
			`PAYPALISHIRING`,
			4,
			`PINALSIGYAHRPI`,
		},
		{
			`PAYPALISHIRING`,
			5,
			`PHASIYIRPLIGAN`,
		},
		{
			"ABC",
			1,
			"ABC",
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

func BenchmarkConvert(b *testing.B)  {
	test := getTestcases()[2]
	b.ResetTimer()
	for i:=0; i < b.N; i++ {
		convert(test.input,test.numRows)
	}
}

func BenchmarkConvert2(b *testing.B)  {
	test := getTestcases()[2]
	b.ResetTimer()
	for i:=0; i < b.N; i++ {
		convert2(test.input,test.numRows)
	}
}
