package Interpreter

type AddNode struct {
	left, right Node
}

func (n *AddNode) Interpreter() int {
	return n.left.Interpreter() + n.right.Interpreter()
}
