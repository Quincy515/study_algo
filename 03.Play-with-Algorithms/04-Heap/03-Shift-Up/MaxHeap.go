package Max_Heap

import (
	"reflect"
)

type MaxHeap struct {
	data     []interface{}
	count    int
	capacity int
}

// 构造一个空堆,可容纳capacity个元素
func NewMaxHeap(capacity int) *MaxHeap {
	data := make([]interface{}, capacity+1)
	return &MaxHeap{
		data:     data,
		count:    0,
		capacity: capacity,
	}
}

// 返回堆中的元素个数
func (h *MaxHeap) Size() int {
	return h.count
}

// 返回一个布尔值,表示堆中是否为空
func (h *MaxHeap) IsEmpty() bool {
	return h.count == 0
}

// 向最大堆中插入一个新的元素 item
func (h *MaxHeap) Insert(item interface{}) {
	if h.count+1 > h.capacity {
		panic("out of range.")
	}
	h.data[h.count+1] = item
	h.count++

	h.shiftUp(h.count)
}

func (h *MaxHeap) shiftUp(k int) {
	// 注意k越界   k节点是否大于其父亲节点(k/2)
	for k > 1 && compare(h.data[k/2], h.data[k]) < 0 {
		swap(h.data, k/2, k) // parent(i)=i/2
		k /= 2
	}
}

// 交换堆中索引为i和j的两个元素
func swap(data []interface{}, i, j int) {
	data[i], data[j] = data[j], data[i]
}

func compare(a, b interface{}) int {
	aType := reflect.TypeOf(a).String()
	bType := reflect.TypeOf(b).String()
	if aType != bType {
		panic("cannot compare different type params")
	}
	switch a.(type) {
	case int:
		if a.(int) > b.(int) {
			return 1
		} else if a.(int) < b.(int) {
			return -1
		} else {
			return 0
		}
	case string:
		if a.(string) > b.(string) {
			return 1
		} else if a.(string) < b.(string) {
			return -1
		} else {
			return 0
		}
	default:
		panic("unsupported type params")
	}
}
