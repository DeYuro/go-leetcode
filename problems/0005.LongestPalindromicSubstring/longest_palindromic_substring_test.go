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

func BenchmarkLongestPalindrome(b *testing.B)  {
	test := getTestcases()[2]
	b.ResetTimer()
	for i:=0; i < b.N; i++ {
		longestPalindrome(test.input)
	}
}

func BenchmarkLongestPalindrome2(b *testing.B)  {
	test := getTestcases()[2]
	b.ResetTimer()
	for i:=0; i < b.N; i++ {
		longestPalindrome2(test.input)
	}
}

func BenchmarkLongestPalindrome3(b *testing.B)  {
	test := getTestcases()[2]
	b.ResetTimer()
	for i:=0; i < b.N; i++ {
		longestPalindrome3(test.input)
	}
}