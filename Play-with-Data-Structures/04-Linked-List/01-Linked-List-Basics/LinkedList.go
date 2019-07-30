package _1_Linked_List_Basics

import "fmt"

type Node struct {
	e    interface{}
	next *Node
}

func NewNode(e interface{}, next *Node) *Node {
	return &Node{e, next}
}

func NewNodeNoN(e interface{}, next *Node) *Node {
	return &Node{e, nil}
}

func NewNodeNo(e interface{}, next *Node) *Node {
	return &Node{nil, nil}
}

func (n *Node) String() string {
	return fmt.Sprint(n.e)
}
