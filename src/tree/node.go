package tree

import "fmt"

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

func (node TreeNode) Print() {
	fmt.Print(node.Value)
}

func (node *TreeNode) SetValue(value int) {
	node.Value = value
}

func (node *TreeNode) CreateNode(value int) {
	return
}

func (node *TreeNode) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}
