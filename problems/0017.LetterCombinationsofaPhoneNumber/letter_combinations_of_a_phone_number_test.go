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
			"7",
			[]string{"p","q","r","s"},
		},
		{
			"27",
			[]string{"ap","aq","ar","as","bp","bq","br","bs","cp","cq","cr","cs"},
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