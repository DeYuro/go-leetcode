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

func BenchmarkLengthOfLongestSubstring(b *testing.B) {
	test := getTestcases()[2]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lengthOfLongestSubstring2(test.input)
	}
}

func BenchmarkLengthOfLongestSubstring2(b *testing.B) {
	test := getTestcases()[2]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lengthOfLongestSubstring2(test.input)
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
