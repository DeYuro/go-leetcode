package problem0242

import (
	"reflect"
	"testing"
)

type test struct {
	s string
	t string
	expected bool
}

func getTestcases() []test {
	return []test{
		{
			"anagram",
			"nagaram",
			true,
		},
		{
			"rat",
			"car",
			false,
		},
	}
}
func TestValidAnagram(t *testing.T) {
	for _, v := range getTestcases() {
		output := isAnagram(v.s,v.t)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}

func TestValidAnagram2(t *testing.T) {
	for _, v := range getTestcases() {
		output := isAnagram2(v.s,v.t)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}

func TestValidAnagram3(t *testing.T) {
	for _, v := range getTestcases() {
		output := isAnagram3(v.s,v.t)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}