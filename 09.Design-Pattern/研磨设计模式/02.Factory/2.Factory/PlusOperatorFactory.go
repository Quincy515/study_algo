package Factory

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
