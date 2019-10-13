package tree

import "fmt"

type Node struct {
	data  interface{}
	left  *Node
	right *Node
}

func NewNode(data interface{}) *Node {
	return &Node{data: data}
}

func (this *Node) String() string {
	return fmt.Sprintf("v: %+v, left: %+v, right: %+v", this.data, this.left, this.right)
}

type BinaryTree struct {
	root *Node
}

func NewBinaryTree(rootV interface{}) *BinaryTree {
	return &BinaryTree{NewNode(rootV)}
}
