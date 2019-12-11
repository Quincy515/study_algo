package _6_Implement_Stack_in_LinkedList

type Stack interface {
	GetSize() int
	IsEmpty() bool
	Push(interface{})
	Pop() interface{}
	Peek() interface{}
}
