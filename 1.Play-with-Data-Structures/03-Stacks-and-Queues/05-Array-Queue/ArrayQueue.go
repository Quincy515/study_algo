package _5_Array_QueueArray

import (
	"bytes"
	"fmt"
)

type ArrayQueue struct {
	array *Array
}

// 构造函数
func NewArrayQueue(capacity int) *ArrayQueue {
	return &ArrayQueue{array: NewArray(capacity)}
}

// 无参数的构造函数
func NewArrayQueueNo() *ArrayQueue {
	return &ArrayQueue{array: NewArrayNP()}
}

func (q *ArrayQueue) GetSize() int {
	return q.array.GetSize()
}

func (q *ArrayQueue) IsEmpty() bool {
	return q.array.IsEmpty()
}

func (q *ArrayQueue) GetCapacity() int {
	return q.array.GetCapacity()
}

// 入队，添加到队尾
func (q *ArrayQueue) Enqueue(e interface{}) {
	q.array.AddLast(e)
}

// 出队，从队首取出
func (q *ArrayQueue) Dequeue() interface{} {
	return q.array.RemoveFirst()
}

// 查看队首元素
func (q *ArrayQueue) GetFront() interface{} {
	return q.array.GetFirst()
}

func (q *ArrayQueue) String() string {
	var buffer bytes.Buffer

	buffer.WriteString("Queue: ")
	buffer.WriteString("front [")
	for i := 0; i < q.array.GetSize(); i++ {
		// fmt.Sprint 将 interface{} 类型转换为字符串
		buffer.WriteString(fmt.Sprint(q.array.Get(i)))
		if i != (q.array.GetSize() - 1) {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("] tail")
	return buffer.String()
}
