package problem13

import (
	"reflect"
	"testing"
)

type test struct {
	input    string
	expected int
}

func getTestcases() []test {
	return []test{
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
		output := romanToInt(v.input)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}