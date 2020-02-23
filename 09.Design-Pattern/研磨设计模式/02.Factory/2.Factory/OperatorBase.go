package Factory

// 数据的抽象 - 可以扩展多种数据类型
// OperatorBase 是 Operator 接口实现的基类，封装公用方法
type OperatorBase struct {
	left, right int
}

// 赋值给左右操作数
func (op *OperatorBase) SetLeft(left int) {
	op.left = left
}
func (op *OperatorBase) SetRight(right int) {
	op.right = right
}
