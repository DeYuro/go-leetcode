package problem49

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type test struct {
	input []string
	expected [][]string
}

func getTestcases() []test {
	return []test{
		{
			[]string{"eat","tea","tan","ate","nat","bat"},
			[][]string{{"bat"},{"nat","tan"},{"ate","eat","tea"}},
		},
		{
			[]string{},
			[][]string{},
		},
		{
			[]string{"a"},
			[][]string{{"a"}},
		},
	}
}

type dummyt struct{}

func (t dummyt) Errorf(string, ...interface{}) {}
//Todo correct check
func TestGroupAnagrams(t *testing.T)  {
	for _,v := range getTestcases() {
		output := groupAnagrams(v.input)
		if !assert.ElementsMatch(dummyt{},output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}
func TestGroupAnagrams2(t *testing.T)  {
	for _,v := range getTestcases() {
		output := groupAnagrams2(v.input)
		if !assert.ElementsMatch(dummyt{},output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}
func TestGroupAnagrams3(t *testing.T)  {
	for _,v := range getTestcases() {
		output := groupAnagrams3(v.input)
		if !assert.ElementsMatch(dummyt{},output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}
func TestGroupAnagrams4(t *testing.T)  {
	for _,v := range getTestcases() {
		output := groupAnagrams4(v.input)
		if !assert.ElementsMatch(dummyt{},output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}
func BenchmarkGroupAnagrams(b *testing.B)  {
	test := getTestcases()[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		groupAnagrams(test.input)
	}
}

func BenchmarkGroupAnagrams2(b *testing.B)  {
	test := getTestcases()[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		groupAnagrams(test.input)
	}
}

func BenchmarkGroupAnagrams3(b *testing.B)  {
	test := getTestcases()[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		groupAnagrams3(test.input)
	}
}

func BenchmarkGroupAnagrams4(b *testing.B)  {
	test := getTestcases()[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		groupAnagrams4(test.input)
	}
}