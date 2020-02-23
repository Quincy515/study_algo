package Prototype

import (
	"fmt"
	"testing"
)

type Type1 struct {
	name string
}

func (t *Type1) Clone() Cloneable {
	tc := *t   // 开辟内存新建变量，存储指针指向的内容
	return &tc // 返回地址 - 深拷贝
	//return t // 浅拷贝
}

type Type2 struct {
	name string
}

func (t *Type2) Clone() Cloneable {
	tc := *t
	return &tc
}

func TestPrototype(t *testing.T) {
	manager := NewPrototypeManager()
	t1 := &Type1{name: "type1"}
	manager.Set("t1", t1)

	t11 := manager.Get("t1")
	t22 := t11.Clone() // 复制
	if t11 == t22 {
		fmt.Println("浅拷贝")
	} else {
		fmt.Println("深拷贝")
	}
}
