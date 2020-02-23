package Interpreter

type MinNode struct {
	left, right Node
}

func (n *MinNode) Interpreter() int {
	return n.left.Interpreter() - n.right.Interpreter()
}
