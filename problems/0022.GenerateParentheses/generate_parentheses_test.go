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
		//{
		//	4,
		//	[]string{"((()))","(()())","(())()","()(())","()()()"},
		//},
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