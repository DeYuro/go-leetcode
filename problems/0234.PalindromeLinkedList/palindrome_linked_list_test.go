package problem234

import (
	"github.com/deyuro/go-leetcode/structures"
	"reflect"
	"testing"
)

type test struct {
	input       *structures.ListNode
	expected bool
}

func getTestcases() []test {
	return []test{
		{
			structures.CreateListNode([]int{1,2,4,2,1}),
			true,
		},
		{
			structures.CreateListNode([]int{1,1,2,1}),
			false,
		},
		{
			structures.CreateListNode([]int{1,2,2,1}),
			true,
		},
		{
			structures.CreateListNode([]int{1,2}),
			false,
		},
	}
}

func TestIsPalindrome(t *testing.T)  {
	for _,v := range getTestcases() {
		output := isPalindrome(v.input)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}

func TestIsPalindrome2(t *testing.T)  {
	for _,v := range getTestcases() {
		output := isPalindrome2(v.input)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}

func BenchmarkIsPalindrome1(b *testing.B)  {
	test := getTestcases()[0]
	b.ResetTimer()
	for i:=0; i < b.N; i++ {
		isPalindrome(test.input)
	}
}

func BenchmarkIsPalindrome2(b *testing.B)  {
	test := getTestcases()[0]
	b.ResetTimer()
	for i:=0; i < b.N; i++ {
		isPalindrome2(test.input)
	}
}