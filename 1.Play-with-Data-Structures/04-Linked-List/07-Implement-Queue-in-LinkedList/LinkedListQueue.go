package _7_Implement_Queue_in_LinkedList

import (
	"bytes"
	"fmt"
)

type node struct {
	e    interface{}
	next *node
}

func (n *node) String() string {
	return fmt.Sprint(n.e)
}

type LinkedListQueue struct {
	head, tail *node
	size       int
}

func NewLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue{}
}

func (lq *LinkedListQueue) GetSize() int {
	return lq.size
}

func (lq *LinkedListQueue) IsEmpty() bool {
	return lq.size == 0
}

func (lq *LinkedListQueue) Enqueue(e interface{}) {
	if lq.tail == nil {
		lq.tail = &node{e: e}
		lq.head = lq.tail
	} else {
		lq.tail.next = &node{e: e}
		lq.tail = lq.tail.next
	}
	lq.size++
}

func (lq *LinkedListQueue) Dequeue() interface{} {
	if lq.IsEmpty() {
		panic("Cannot dequeue from an empty queue.")
	}
	retNode := lq.head
	lq.head = lq.head.next
	retNode.next = nil
	if retNode == nil {
		lq.tail = nil
	}
	lq.size--
	return retNode.e
}

func (lq *LinkedListQueue) GetFront() interface{} {
	if lq.IsEmpty() {
		panic("Queue is Empty.")
	}
	return lq.head.e
}

func (lq *LinkedListQueue) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("Queue: front ")

	cur := lq.head
	for cur != nil {
		buffer.WriteString(fmt.Sprintf("%v -> ", cur.e))
		cur = cur.next
	}
	buffer.WriteString("NULL tail")
	return buffer.String()
}
