package problem94

import (
	"github.com/deyuro/go-leetcode/structures"
	"reflect"
	"testing"
)

type test struct {
	input *structures.TreeNode
	expected []int
}

func getTestcases() []test {
	return []test{
		{
			&structures.TreeNode{
				Val:   1,
				Left:  nil,
				Right: &structures.TreeNode{
					Val:   2,
					Left:  &structures.TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
			},
			[]int{1,3,2},
		},
		{
			nil,
			[]int{},
		},
		{
			&structures.TreeNode{
				Val:   1,
				Left:  &structures.TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
			[]int{2,1},
		},
		{
			&structures.TreeNode{
				Val:   1,
				Left: nil,
				Right:  &structures.TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
			},
			[]int{1,2},
		},
		{
			&structures.TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			[]int{1},
		},
	}
}

func TestInorderTraversal(t *testing.T) {
	for _, v := range getTestcases() {
		output := inTraversal(v.input)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}

func BenchmarkInorderTraversal(b *testing.B) {
	test := getTestcases()[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		inTraversal(test.input)
	}
}