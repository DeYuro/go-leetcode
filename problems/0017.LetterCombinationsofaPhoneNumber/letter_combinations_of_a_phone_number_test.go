package problem17

import (
	"reflect"
	"testing"
)

type test struct {
	input    string
	expected []string
}

func getTestcases() []test {
	return []test{
		{
			"23",
			[]string{"ad","ae","af","bd","be","bf","cd","ce","cf"},
		},
		{
			"",
			[]string{},
		},
		{
			"2",
			[]string{"a","b","c"},
		},
	}
}

func TestLetterCombinations(t *testing.T) {
	for _, v := range getTestcases() {
		output := letterCombinations(v.input)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}