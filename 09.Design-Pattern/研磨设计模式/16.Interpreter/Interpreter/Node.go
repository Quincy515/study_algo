package Interpreter

// 1+2-3 加法减法数据
// Node: 节点,返回数据
type Node interface {
	Interpreter() int
}

type ValNode struct {
	val int
}

func (n *ValNode) Interpreter() int {
	return n.val
}
