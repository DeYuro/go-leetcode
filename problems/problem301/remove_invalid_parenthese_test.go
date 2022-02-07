package problem301

import (
	"testing"
)

type test struct {
	input string
	expected []string
}

func getTestcases() []test {
	return []test{
		{
			"()())()",
			[]string{"(())()","()()()"},
		},
		{
			"(a)())()",
			[]string{"(())()","()()()"},
		},
	}
}
func TestRemoveInvalidParentheses(t *testing.T) {
	for _, v := range getTestcases() {
		output := removeInvalidParentheses(v.input)
		if len(output) != len (v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}

func BenchmarkRemoveInvalidParentheses(b *testing.B) {
	test := getTestcases()[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		removeInvalidParentheses(test.input)
	}
}