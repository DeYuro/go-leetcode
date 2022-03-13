package problem58

import (
	"reflect"
	"testing"
)

type test struct {
	input    string
	expected int
}

func getTestcases() []test {
	return []test{
		{
			"Hello World",
			5,
		},
		{
			"   fly me   to   the moon  ",
			4,
		},
		{
			"luffy is still joyboy",
			6,
		},
	}
}

func TestLengthOfLastWord(t *testing.T) {
	for _, v := range getTestcases() {
		output := lengthOfLastWord(v.input)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}

func TestLengthOfLastWord2(t *testing.T) {
	for _, v := range getTestcases() {
		output := lengthOfLastWord2(v.input)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}

func BenchmarkLengthOfLastWord(b *testing.B) {
	test := getTestcases()[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lengthOfLastWord(test.input)
	}
}

func BenchmarkLengthOfLastWord2(b *testing.B) {
	test := getTestcases()[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lengthOfLastWord2(test.input)
	}
}