package Implement

// A X B
// X 代表操作 A、B 代表操作数
// X=+ A+B X=- A-B X=% A%B
// Operator 是被封装的实际运行类的接口
type Operator interface {
	SetLeft(int)
	SetRight(int)
	Result() int
}

// OperatorFactory 是工厂接口
type OperatorFactory interface {
	Create() Operator
}

// 数据的抽象 - 可以扩展多种数据类型
//type OperatorBaseFloat struct {
//	left, right float64
//}

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

// 操作的抽象
// PlusOperatorFactory 是 PlusOperator 的工厂类
type PlusOperatorFactory struct{}

func (PlusOperatorFactory) Create() Operator {
	return &PlusOperator{&OperatorBase{}}
}

// 操作类中包含操作数
// PlusOperator 是 Operator 的实际加法实现
type PlusOperator struct {
	*OperatorBase
}

// Result 实际运行,获取结果
func (o PlusOperator) Result() int {
	return o.right + o.left
}

// 操作的抽象
// SubOperatorFactory 是 SubOperator 的工厂类
type SubOperatorFactory struct{}

// 操作类中包含操作数
// SubOperator 是 Operator 的实际减法实现
type SubOperator struct {
	*OperatorBase
}

// Result 实际运行,获取结果
func (o SubOperator) Result() int {
	return o.left - o.right
}

func (SubOperatorFactory) Create() Operator {
	return &SubOperator{&OperatorBase{}}
}
