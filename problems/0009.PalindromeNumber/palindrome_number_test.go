package problem9

import (
	"reflect"
	"testing"
)

type test struct {
	input    int
	expected bool
}

func TestIsPalindrome(t *testing.T) {
	for _, v := range getTestcases() {
		output := isPalindrome(v.input)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}

func TestIsPalindrome2(t *testing.T) {
	for _, v := range getTestcases() {
		output := isPalindrome2(v.input)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	test := getTestcases()[2]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isPalindrome(test.input)
	}
}

func BenchmarkIsPalindrome2(b *testing.B) {
	test := getTestcases()[2]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isPalindrome2(test.input)
	}
}

func getTestcases() []test {
	return []test{
		{
			121,
			true,
		},
		{
			-121,
			false,
		}, {
			10,
			false,
		}, {
			-101,
			false,
		},
		{
			20,
			false,
		}, {
			123,
			false,
		},
	}
}
