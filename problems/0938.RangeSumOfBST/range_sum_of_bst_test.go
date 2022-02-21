package problem938

import (
	"github.com/deyuro/go-leetcode/structures"
	"reflect"
	"testing"
)

type test struct {
	input     *structures.TreeNode
	low, high int
	expected  int
}

func getTestcases() []test {
	return []test{
		{
			&structures.TreeNode{
				Val: 10,
				Left: &structures.TreeNode{
					Val: 5,
					Left: &structures.TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
					Right: &structures.TreeNode{
						Val:   7,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &structures.TreeNode{
					Val:  15,
					Left: nil,
					Right: &structures.TreeNode{
						Val:   18,
						Left:  nil,
						Right: nil,
					},
				},
			},
			7,
			15,
			32,
		},
	}
}

func TestIsRangeSumBST(t *testing.T) {
	for _, v := range getTestcases() {
		output := rangeSumBST(v.input,v.low,v.high)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}

func BenchmarkRangeSumBST(b *testing.B) {
	test := getTestcases()[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rangeSumBST(test.input,test.low,test.high)
	}
}