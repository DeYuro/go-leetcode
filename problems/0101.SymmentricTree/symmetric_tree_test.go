package problem100

import (
	"github.com/deyuro/go-leetcode/structures"
	"reflect"
	"testing"
)

type test struct {
	input *structures.TreeNode
	expected bool
}

func getTestcases() []test {
	return []test{
		{
			&structures.TreeNode{
				Val: 1,
				Left: &structures.TreeNode{
					Val:   2,
					Left:  &structures.TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
					Right: &structures.TreeNode{
						Val:   4,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &structures.TreeNode{
					Val:   2,
					Left:  &structures.TreeNode{
						Val:   4,
						Left:  nil,
						Right: nil,
					},
					Right: &structures.TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
				},
			},
			true,
		},
		{
			&structures.TreeNode{
				Val: 1,
				Left: &structures.TreeNode{
					Val:   2,
					Left:  nil,
					Right: &structures.TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &structures.TreeNode{
					Val:   2,
					Left:  nil,
					Right: &structures.TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
				},
			},
			false,
		},
	}
}

func TestIsSymmetric(t *testing.T) {
	for _, v := range getTestcases() {
		output := isSymmetric(v.input)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}

func BenchmarkIsSymmetric(b *testing.B) {
	test := getTestcases()[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isSymmetric(test.input)
	}
}