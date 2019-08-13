package _7_Implementation_of_Loop_Queue

import (
	"bytes"
	"fmt"
)

type LoopQueue struct {
	data        []interface{}
	front, tail int
	size        int
}

func NewLoopQueue(capacity int) *LoopQueue {
	loop := &LoopQueue{}
	// 循环队列存在一个空位，所以参数为10，就需要开辟11个位置
	loop.data = make([]interface{}, capacity+1)
	loop.front, loop.tail, loop.size = 0, 0, 0
	return loop
}

func NewLoopQueueNo() *LoopQueue {
	loop := &LoopQueue{}
	loop.data = make([]interface{}, 10)
	loop.front, loop.tail, loop.size = 0, 0, 0
	return loop
}

func (l *LoopQueue) GetCapacity() int {
	// 与上对于，需要删除空位
	return len(l.data) - 1
}

func (l *LoopQueue) IsEmpty() bool {
	return l.front == l.tail
}

func (l *LoopQueue) GetSize() int {
	return l.size
}

// 循环队列入队
func (l *LoopQueue) Enqueue(e interface{}) {
	if (l.tail+1)%len(l.data) == l.front {
		// 判断队列 满
		l.resize(l.GetCapacity() * 2)
	}
	l.data[l.tail] = e
	l.tail = (l.tail + 1) % len(l.data)
	l.size++
}

// 循环队列出队
func (l *LoopQueue) Dequeue() interface{} {
	if l.IsEmpty() {
		panic("Cannot dequeue from an empty queue.")
	}
	ret := l.data[l.front]
	l.data[l.front] = nil
	l.front = (l.front + 1) % len(l.data)
	l.size--

	if l.size == l.GetCapacity()/4 && l.GetCapacity()/2 != 0 {
		l.resize(l.GetCapacity() / 2)
	}

	return ret
}

func (l *LoopQueue) GetFront() interface{} {
	if l.IsEmpty() {
		panic("Queue is empty.")
	}
	return l.data[l.front]
}

func (l *LoopQueue) resize(newCapacity int) {
	newData := make([]interface{}, newCapacity+1) // 开辟新的空间
	// data所有数据放入到newData中
	for i := 0; i < l.size; i++ {
		newData[i] = l.data[(i+l.front)%len(l.data)]
	}
	l.data = newData
	l.front = 0
	l.tail = l.size
}

func (l *LoopQueue) String() string {
	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("Queue: size = %d, capacity = %d\n", l.size, l.GetCapacity()))
	buffer.WriteString("front [")
	for i := l.front; i != l.tail; i = (i + 1) % len(l.data) {
		// fmt.Sprint 将 interface{} 类型转换为字符串
		buffer.WriteString(fmt.Sprint(l.data[i]))
		if (i+1)%len(l.data) != l.tail {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("] tail")
	return buffer.String()
}
