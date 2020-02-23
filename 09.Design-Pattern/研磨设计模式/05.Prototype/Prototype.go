package Prototype

// Cloneable: 原型对象需要实现的接口
type Cloneable interface {
	Clone() Cloneable
}

// PrototypeManager: 原型对象的类
type PrototypeManager struct {
	prototypes map[string]Cloneable
}

// NewPrototypeManager: 构造初始化
func NewPrototypeManager() *PrototypeManager {
	return &PrototypeManager{make(map[string]Cloneable)}
}

// Get: 设置
func (p *PrototypeManager) Get(name string) Cloneable {
	return p.prototypes[name]
}

// Set: 获取
func (p *PrototypeManager) Set(name string, prototype Cloneable) {
	p.prototypes[name] = prototype
}
