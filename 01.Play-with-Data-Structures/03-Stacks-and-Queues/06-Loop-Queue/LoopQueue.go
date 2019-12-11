package _6_Loop_Queue

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
