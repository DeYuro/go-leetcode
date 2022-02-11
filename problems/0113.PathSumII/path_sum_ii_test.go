package problem113

import (
	"github.com/deyuro/go-leetcode/structures"
	"reflect"
	"testing"
)

type test struct {
	input *structures.TreeNode
	target int
	expected [][]int
}

func getTestcases() []test {
	return []test{
		{
			&structures.TreeNode{
				Val:   5,
				Left:  &structures.TreeNode{
					Val:   4,
					Left:  &structures.TreeNode{
						Val:   11,
						Left:  &structures.TreeNode{
							Val:   7,
							Left:  nil,
							Right: nil,
						},
						Right: &structures.TreeNode{
							Val:   2,
							Left:  nil,
							Right: nil,
						},
					},
					Right: nil,
				},
				Right: &structures.TreeNode{
					Val:   8,
					Left:  &structures.TreeNode{
						Val:   13,
						Left:  nil,
						Right: nil,
					},
					Right: &structures.TreeNode{
						Val:   4,
						Left:  &structures.TreeNode{
							Val:   5,
							Left:  nil,
							Right: nil,
						},
						Right: &structures.TreeNode{
							Val:   1,
							Left:  nil,
							Right: nil,
						},
					},
				},
			},
			22,
			[][]int{{5,4,11,2},{5,8,4,5}},
		},
	}
}

func TestPathSum(t *testing.T) {
	for _, v := range getTestcases() {
		output := pathSum(v.input, v.target)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}

func BenchmarkIsBalanced(b *testing.B) {
	test := getTestcases()[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pathSum(test.input, test.target)
	}
}
