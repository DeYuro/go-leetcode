package problem5

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
			`babad`,
			`bab`,
		},
		{
			`cbbd`,
			`bb`,
		},
		{
			`a`,
			`a`,
		},
		{
			`ac`,
			`a`,
		},
	}
}

func TestLongestPalindrome(t *testing.T) {
	for _, v := range getTestcases() {
		output := longestPalindrome(v.input)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}