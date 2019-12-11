package Max_Heap

type MaxHeap struct {
	data  []int
	count int
}

// 构造一个空堆,可容纳capacity个元素
func NewMaxHeap(capacity int) *MaxHeap {
	data := make([]int, capacity+1)
	return &MaxHeap{
		data:  data,
		count: 0,
	}
}

// 返回堆中的元素个数
func (h *MaxHeap) Size() int {
	return h.count
}

// 返回一个布尔值,表示堆中是否为空
func (h MaxHeap) IsEmpty() bool {
	return h.count == 0
}
