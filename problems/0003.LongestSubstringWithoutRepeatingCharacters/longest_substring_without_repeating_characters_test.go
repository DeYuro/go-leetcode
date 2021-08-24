package problem3

import (
	"reflect"
	"testing"
)

type test struct {
	input    string
	expected int
}

func TestLengthOfLongestSubstring(t *testing.T)  {
	for _, v := range getTestcases() {
		output := lengthOfLongestSubstring(v.input)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}
func TestLengthOfLongestSubstring2(t *testing.T)  {
	for _, v := range getTestcases() {
		output := lengthOfLongestSubstring2(v.input)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}
func getTestcases() []test {
	return []test{
		{
			`abcabcbb`,
			3,
		},
		{
			`bbbbb`,
			1,
		},
		{
			`pwwkew`,
			3,
		},
		{
			``,
			0,
		},
		{
			`aab`,
			2,
		},
		{
			`dvdf`,
			3,
		},
	}
}
