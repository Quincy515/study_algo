package Factory

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
