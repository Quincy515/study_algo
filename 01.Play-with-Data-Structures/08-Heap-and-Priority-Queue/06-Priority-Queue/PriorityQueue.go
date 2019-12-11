package _6_Priority_Queue

type PriorityQueue struct {
	maxHeap *MaxHeap
}

func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{NewMaxHeap(20)}
}

func (pq *PriorityQueue) GetSize() int {
	return pq.maxHeap.Size()
}

func (pq *PriorityQueue) IsEmpty() bool {
	return pq.maxHeap.IsEmpty()
}

func (pq *PriorityQueue) GerFront() interface{} {
	return pq.maxHeap.FindMax()
}

func (pq *PriorityQueue) Enqueue(e interface{}) {
	pq.maxHeap.Add(e)
}

func (pq *PriorityQueue) Dequeue() interface{} {
	return pq.maxHeap.ExtractMax()
}
