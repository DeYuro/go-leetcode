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

func TestLongestCommonPrefix2(t *testing.T) {
	for _, v := range getTestcases() {
		output := longestCommonPrefix2(v.input)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}

func BenchmarkLongestCommonPrefix(b *testing.B)  {
	test := getTestcases()[2]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		longestCommonPrefix(test.input)
	}
}

func BenchmarkLongestCommonPrefix2(b *testing.B)  {
	test := getTestcases()[2]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		longestCommonPrefix2(test.input)
	}
}