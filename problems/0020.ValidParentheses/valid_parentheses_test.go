package problem20

import (
	"reflect"
	"testing"
)

type test struct {
	input    string
	expected bool
}

func TestIsValid(t *testing.T) {
	for _, v := range getTestcases() {
		output := isValid(v.input)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}

func TestIsValid2(t *testing.T) {
	for _, v := range getTestcases() {
		output := isValid2(v.input)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}

func BenchmarkIsValid(b *testing.B) {
	test := getTestcases()[2]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isValid(test.input)
	}
}

func BenchmarkIsValid2(b *testing.B) {
	test := getTestcases()[2]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isValid2(test.input)
	}
}
func getTestcases() []test {
	return []test{
		{
			`()[]{}`,
			true,
		},
		{
			`()`,
			true,
		}, {
			`(]`,
			false,
		}, {
			`([)]`,
			false,
		},
		{
			`{[]}`,
			true,
		},
		{
			``,
			true,
		},
		{
			`][`,
			false,
		},
	}
}