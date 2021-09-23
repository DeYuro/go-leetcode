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
			[]string{"ad","bd","cd","ae","be","ce","af","bf","cf"},
		},
		{
			"7",
			[]string{"p","q","r","s"},
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


func BenchmarkIsValid(b *testing.B) {
	test := getTestcases()[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		letterCombinations(test.input)
	}
}