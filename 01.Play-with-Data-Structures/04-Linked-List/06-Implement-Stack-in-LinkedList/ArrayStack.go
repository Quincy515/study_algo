package _6_Implement_Stack_in_LinkedList

import (
	"bytes"
	"fmt"
)

type ArrayStack struct {
	array *Array
}

// 构造函数
func NewArrayStack(capacity int) *ArrayStack {
	return &ArrayStack{array: NewArray(capacity)}
}

// 没有参数的构造函数
func NewArrayStackNo() *ArrayStack {
	return &ArrayStack{array: NewArrayNP()}
}

func (a *ArrayStack) GetSize() int {
	return a.array.GetSize()
}

func (a *ArrayStack) IsEmpty() bool {
	return a.array.IsEmpty()
}

func (a *ArrayStack) GetCapacity() int {
	return a.array.GetCapacity()
}

// 压入栈顶元素
func (a *ArrayStack) Push(e interface{}) {
	a.array.AddLast(e)
}

// 弹出栈顶元素
func (a *ArrayStack) Pop() interface{} {
	return a.array.RemoveLast()
}

// 查看栈顶元素
func (a *ArrayStack) Peek() interface{} {
	return a.array.GetLast()
}

func (a *ArrayStack) String() string {
	var buffer bytes.Buffer

	buffer.WriteString("Stack: ")
	buffer.WriteString("[")
	for i := 0; i < a.array.GetSize(); i++ {
		buffer.WriteString(fmt.Sprint(a.array.Get(i)))
		if i != a.array.GetSize()-1 {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("] top")

	return buffer.String()
}
