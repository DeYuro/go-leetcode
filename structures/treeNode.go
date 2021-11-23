package structures

type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}

type Bst struct {
	root *TreeNode
}

func (b *Bst) Insert(value int) {
	b.insertRec(b.root, value)
}

func (b *Bst) insertRec(node *TreeNode, value int) *TreeNode {
	if b.root == nil {
		b.root = &TreeNode{
			Val: value,
		}
		return b.root
	}
	if node == nil {
		return &TreeNode{Val: value}
	}
	if value <= node.Val {
		node.Left = b.insertRec(node.Left, value)
	}
	if value > node.Val {
		node.Right = b.insertRec(node.Right, value)
	}
	return node
}

func BstFromSlice(s []int) *TreeNode {
	bst := Bst{}

	for _, v := range s {
		bst.Insert(v)
	}

	return bst.root
}