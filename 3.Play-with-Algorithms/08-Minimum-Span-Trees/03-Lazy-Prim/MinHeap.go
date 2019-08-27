package Lazy_Prim

type MinHeap struct {
	data     []interface{}
	count    int
	capacity int
}

func (h *MinHeap) shiftUp(k int) {
	// 注意k越界   k节点是否小于其父亲节点(k/2)
	for k > 1 && compare(h.data[k/2], h.data[k]) > 0 {
		swap(h.data, k/2, k) // parent(i)=i/2
		k /= 2
	}
}

func (h MinHeap) shiftDown(k int) {
	for 2*k <= h.count { // 完全二叉树中有左孩子，就有孩子
		j := 2 * k // 在此轮循环中,data[k]和data[j]交换位置
		// 判断是否有右孩子        右孩子和左孩子比较大小
		if j+1 <= h.count && compare(h.data[j+1], h.data[j]) < 0 {
			j += 1 // 变更j为右孩子
		}
		if compare(h.data[k], h.data[j]) <= 0 {
			break // 满足二叉堆定义, 任意节点不大于其父节点
		}
		swap(h.data, k, j) // 不满足二叉堆定义,就交换k和左右孩子中较大的孩子
		k = j              // 继续进行下浮操作,知道满足二叉堆定义
	}
}

// 构造一个空堆,可容纳capacity个元素
func NewMinHeap(capacity int) *MinHeap {
	data := make([]interface{}, capacity+1)
	return &MinHeap{
		data:     data,
		count:    0,
		capacity: capacity,
	}
}

// 构造函数,通过一个给定数组创建一个最大堆
// 该构造堆的过程,时间复杂度为O(n)
func NewHeapify(arr []interface{}, n int) *MinHeap {
	data := make([]interface{}, n+1)
	capacity := n
	for i := 0; i < n; i++ {
		data[i+1] = arr[i]
	}
	count := n
	maxHeap := &MinHeap{data, capacity, count}

	for i := count / 2; i >= 1; i-- {
		maxHeap.shiftDown(i)
	}
	return maxHeap
}

// 返回堆中的元素个数
func (h *MinHeap) Size() int {
	return h.count
}

// 返回一个布尔值,表示堆中是否为空
func (h *MinHeap) IsEmpty() bool {
	return h.count == 0
}

// 向最大堆中插入一个新的元素 item
func (h *MinHeap) Insert(item interface{}) {
	if h.count+1 > h.capacity {
		panic("out of range.")
	}
	h.data[h.count+1] = item
	h.count++

	h.shiftUp(h.count)
}

// 从最大堆中取出堆顶元素,即堆中所存储的最大数据
func (h *MinHeap) ExtractMin() interface{} {
	if h.count <= 0 {
		panic("堆中没有元素!")
	}
	ret := h.data[1]
	swap(h.data, 1, h.count) // 最后一个元素和第一个元素交换
	h.count--
	h.shiftDown(1)
	return ret
}

// 交换堆中索引为i和j的两个元素
func swap(data []interface{}, i, j int) {
	data[i], data[j] = data[j], data[i]
}
