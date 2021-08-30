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