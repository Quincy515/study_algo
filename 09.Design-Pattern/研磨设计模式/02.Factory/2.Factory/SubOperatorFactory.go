package Factory

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
