package problem12

import (
	"reflect"
	"testing"
)

type test struct {
	expected string
	input int
}

func getTestcases() []test {
	return []test{
		{
			"LXXX",
			80,
		},
		{
			"III",
			3,
		},
		{
			"IV",
			4,
		}, {
			"IX",
			9,
		},
		{
			"LVIII",
			58,
		},
		{
			"MCMXCIV",
			1994,
		},
		{
			"MDCCCLXXXIV",
			1884,
		},
	}
}

func TestRomanToInt(t *testing.T) {
	for _, v := range getTestcases() {
		output := intToRoman(v.input)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}
func TestRomanToInt2(t *testing.T) {
	for _, v := range getTestcases() {
		output := intToRoman2(v.input)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}
func BenchmarkRomanToInt(b *testing.B)  {
	test := getTestcases()[5]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		intToRoman(test.input)
	}
}

func BenchmarkRomanToInt2(b *testing.B)  {
	test := getTestcases()[5]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		intToRoman2(test.input)
	}
}