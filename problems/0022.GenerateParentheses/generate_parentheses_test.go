package problem22

import (
	"reflect"
	"testing"
)

type test struct {
	input    int
	expected []string
}


func getTestcases() []test {
	return []test{
		{
			3,
			[]string{"((()))","(()())","(())()","()(())","()()()"},
		},
		{
			2,
			[]string{"(())","()()"},
		},
		{
			1,
			[]string{"()"},
		},
	}
}

func TestGenerateParenthesis(t *testing.T) {
	for _, v := range getTestcases() {
		output := generateParenthesis(v.input)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}

func BenchmarkGenerateParenthesis(b *testing.B) {
	test := getTestcases()[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		generateParenthesis(test.input)
	}
}