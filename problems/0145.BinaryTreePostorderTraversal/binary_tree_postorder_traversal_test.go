package problem145

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
			[]int{3,2,1},
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
			[]int{2,1},
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

func TestPreorderTraversal(t *testing.T) {
	for _, v := range getTestcases() {
		output := postorderTraversal(v.input)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}

func BenchmarkPreorderTraversal(b *testing.B) {
	test := getTestcases()[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		postorderTraversal(test.input)
	}
}