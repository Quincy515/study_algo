package Composite

import "fmt"

type Leaf struct {
	component
}

// NewLeaf: 开辟一个叶子
func NewLeaf() *Leaf {
	return &Leaf{}
}

func (l *Leaf) Print(pre string) {
	fmt.Printf("%s-%s\n", pre, l.Name())
}

type Composite struct {
	component
	childs []Component // 叶子的集合
}

// NewComposite: 创建一个组合结构体
func NewComposite() *Composite {
	return &Composite{
		childs: make([]Component, 0),
	}
}

func (c *Composite) AddChild(child Component) {
	child.SetParent(c)
	c.childs = append(c.childs, child) // 加入孩子节点
}

// 打印显示每一个节点
func (c *Composite) Print(pre string) {
	fmt.Printf("%s+%s\n", pre, c.Name())
	pre += " "
	for _, comp := range c.childs {
		comp.Print(pre)
	}
}
