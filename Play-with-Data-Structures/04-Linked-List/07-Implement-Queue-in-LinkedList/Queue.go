package _7_Implement_Queue_in_LinkedList

type Queue interface {
	GetSize() int
	IsEmpty() bool
	Enqueue(interface{})
	Dequeue() interface{}
	GetFront() interface{}
}
