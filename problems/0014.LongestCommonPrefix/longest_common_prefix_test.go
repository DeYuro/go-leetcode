package problem14

import (
	"reflect"
	"testing"
)

type test struct {
	input    []string
	expected string
}

func getTestcases() []test {
	return []test{
		{
		[]string{"flower","flow","flight"},
			"fl",
		},
		{
			[]string{"dog","racecar","car"},
			"",
		},
	}
}

func TestLongestCommonPrefix(t *testing.T) {
	for _, v := range getTestcases() {
		output := longestCommonPrefix(v.input)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}
