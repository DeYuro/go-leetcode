package ways

import (
	"reflect"
	"testing"
)

type test struct {
	n int
	m int
	expected int
}

func getTestcases() []test {
	return []test{
		{

			3,
			2,
			3,
		},
	}
}

func TestPaths(t *testing.T) {
	for _,v := range getTestcases() {
		output := paths(v.n,v.m)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}

func TestPathsDynamic(t *testing.T) {
	for _,v := range getTestcases() {
		output := pathsDynamic(v.n,v.m)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}