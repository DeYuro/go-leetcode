package problem438

import (
	"reflect"
	"testing"
)

type test struct {
	s    string
	p   string
	expected []int
}

func getTestcases() []test {
	return []test{
		{
			"cbaebabacd",
			"abc",
			[]int{0,6},
		},
		{
			"abab",
			"ab",
			[]int{0,1,2},
		},
	}
}

func TestFindAnagrams(t *testing.T) {
	for _, v := range getTestcases() {
		output := findAnagrams(v.s, v.p)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}
func TestFindAnagrams2(t *testing.T) {
	for _, v := range getTestcases() {
		output := findAnagrams2(v.s, v.p)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}
func BenchmarkFindAnagrams(b *testing.B) {
	test := getTestcases()[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findAnagrams(test.s, test.p)
	}
}
func BenchmarkFindAnagrams2(b *testing.B) {
	test := getTestcases()[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findAnagrams2(test.s, test.p)
	}
}
