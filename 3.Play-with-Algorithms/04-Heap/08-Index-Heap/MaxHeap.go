package Index_Heap

type IndexMaxHeap struct {
	data     []int // 最大索引堆中的数据
	indexes  []int // 最大索引堆中的索引
	count    int
	capacity int
}

// 构造一个空堆,可容纳capacity个元素
func NewIndexMaxHeap(capacity int) *IndexMaxHeap {
	data := make([]int, capacity+1)
	indexes := make([]int, capacity+1)
	return &IndexMaxHeap{
		data:     data,
		indexes:  indexes,
		count:    0,
		capacity: capacity,
	}
}

// 返回堆中的元素个数
func (h *IndexMaxHeap) Size() int {
	return h.count
}

// 返回一个布尔值,表示堆中是否为空
func (h *IndexMaxHeap) IsEmpty() bool {
	return h.count == 0
}

// 向最大堆中插入一个新的元素 item和元素的所有 i
// 传入的i对用户而言,是从0索引的
func (h *IndexMaxHeap) Insert(item int, i int) {
	if h.count+1 > h.capacity || i+1 < 1 || i+1 > h.capacity {
		panic("out of range.")
	}
	i += 1
	h.data[i] = item
	h.indexes[h.count+1] = i
	h.count++

	h.shiftUp(h.count)
}

// 从最大索引堆中取出堆顶元素,即堆中所存储的最大数据
func (h *IndexMaxHeap) ExtractMax() int {
	if h.count <= 0 {
		panic("堆中没有元素!")
	}
	ret := h.data[h.indexes[1]]
	swap(h.indexes, 1, h.count) // 最后一个元素和第一个元素交换
	h.count--
	h.shiftDown(1)
	return ret
}

// 从最大索引堆中取出堆顶元素的索引
func (h *IndexMaxHeap) ExtractMaxIndex() int {
	if h.count <= 0 {
		panic("堆中没有元素!")
	}
	ret := h.indexes[1] - 1
	swap(h.indexes, 1, h.count)
	h.count--
	h.shiftDown(1)
	return ret
}

// 获取最大索引堆中的堆顶元素
func (h *IndexMaxHeap) GetMax() int {
	if h.count <= 0 {
		panic("堆中没有元素!")
	}
	return h.data[h.indexes[1]]
}

// 获取最大索引堆中堆顶元素的索引
func (h *IndexMaxHeap) GetMaxIndex() int {
	if h.count <= 0 {
		panic("堆中没有元素!")
	}
	return h.indexes[1] - 1
}

// 获取最大索引堆中索引为i的元素
func (h *IndexMaxHeap) GetItem(i int) int {
	if i+1 < 1 || i+1 > h.capacity {
		panic("索引越界!")
	}
	return h.data[i+1]
}

// 将最大索引堆中索引为i的元素修改为newItem
func (h *IndexMaxHeap) change(i int, newItem int) {
	i += 1
	h.data[i] = newItem

	// 找到indexes[j]=i,j表示data[i]在堆中的位置
	// 之后shiftUp(j),再shiftDown(j)-遍历O(n)shiftUp为O(1)
	for j := 1; j <= h.count; j++ {
		if h.indexes[j] == i {
			h.shiftUp(j)
			h.shiftDown(j)
			return
		}
	}

}

// 索引堆中,数据之间的比较根据data的大小进行比较,但实际操作的是所有
func (h *IndexMaxHeap) shiftUp(k int) {
	// 注意k越界   k节点是否大于其父亲节点(k/2)
	for k > 1 && (h.data[h.indexes[k/2]] < h.data[h.indexes[k]]) {
		swap(h.indexes, k/2, k) // parent(i)=i/2
		k /= 2
	}
}

// 索引堆中,数据之间的比较根据data的大小进行比较,但实际操作的是索引
func (h IndexMaxHeap) shiftDown(k int) {
	for 2*k <= h.count { // 完全二叉树中有左孩子，就有孩子
		j := 2 * k // 在此轮循环中,data[k]和data[j]交换位置
		// 判断是否有右孩子        右孩子和左孩子比较大小
		if j+1 <= h.count && h.data[h.indexes[j+1]] > h.data[h.indexes[j]] {
			j += 1 // 变更j为右孩子
		}
		if h.data[h.indexes[k]] >= h.data[h.indexes[j]] {
			break // 满足二叉堆定义, 任意节点不大于其父节点
		}
		swap(h.indexes, k, j) // 不满足二叉堆定义,就交换k和左右孩子中较大的孩子
		k = j                 // 继续进行下浮操作,知道满足二叉堆定义
	}
}

// 交换堆中索引为i和j的两个元素
func swap(data []int, i, j int) {
	data[i], data[j] = data[j], data[i]
}
