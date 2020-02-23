package Adapter

// 适用于版本接口升级
// Target: 是适配的目标接口
type Target interface {
	Request() string
}

// Adaptee: 是被适配的目标接口
type Adaptee interface {
	SpecficRequest() string
}

// NewAdaptee: 是被适配接口的工厂函数
func NewAdaptee() Adaptee {
	return &adapteeImpl{}
}

// adapteeImpl: 是被适配的目标类
type adapteeImpl struct{}

// SpecficRequest: 是目标类的一个方法
func (*adapteeImpl) SpecficRequest() string {
	return "adaptee method SpecficRequest"
}

// adapter: 是转换 Adaptee 为 Target 接口的适配器
type adapter struct {
	Adaptee
}

// Request: 实现 Target 接口
func (a *adapter) Request() string {
	return a.SpecficRequest()
}

// NewAdapter: 是Adapter的工厂函数
func NewAdapter(adaptee Adaptee) Target {
	return &adapter{Adaptee: adaptee}
}
