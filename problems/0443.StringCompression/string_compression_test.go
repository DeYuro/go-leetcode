package problem443

import (
	"reflect"
	"testing"
)

type test struct {
	input    []byte
	expected []byte
}

func getTestcases() []test {
	return []test{
		{
			[]byte{'a','a','b','b','c','c','c'},
			[]byte{'a','2','b','2','c','3'},
		},
		{
			[]byte{'a'},
			[]byte{'a'},
		},
		{
			[]byte{'a','b','b','b','b','b','b','b','b','b','b','b','b'},
			[]byte{'a','b','1','2'},
		},
	}
}


func TestCompress(t *testing.T) {
	for _, v := range getTestcases() {
		output := compress(v.input)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}
