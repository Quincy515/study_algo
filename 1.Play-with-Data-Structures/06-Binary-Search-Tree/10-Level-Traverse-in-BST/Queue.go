package _0_Level_Traverse_in_BST

type Queue interface {
	GetSize() int
	IsEmpty() bool
	Enqueue(interface{})
	Dequeue() interface{}
	GetFront() interface{}
}
