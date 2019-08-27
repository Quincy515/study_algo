package Kruskal

type IndexMinHeap struct {
	data     []interface{} // 最小索引堆中的数据
	indexes  []int         // 最小索引堆中的索引,indexes[x]=i表示索引i在x的位置
	reverse  []int         // 最小索引堆的反向索引,reverse[i]=x表示索引i在x的位置
	count    int
	capacity int
}

// 索引堆中,数据之间的比较根据data的大小进行比较,但实际操作的是所有
func (h *IndexMinHeap) shiftUp(k int) {
	// 注意k越界   k节点是否小于其父亲节点(k/2)
	for k > 1 && compare(h.data[h.indexes[k/2]], h.data[h.indexes[k]]) > 0 {
		swapIndex(h.indexes, k/2, k) // parent(i)=i/2
		h.reverse[h.indexes[k/2]] = k / 2
		h.reverse[h.indexes[k]] = k
		k /= 2
	}
}

// 索引堆中,数据之间的比较根据data的大小进行比较,但实际操作的是索引
func (h IndexMinHeap) shiftDown(k int) {
	for 2*k <= h.count { // 完全二叉树中有左孩子，就有孩子
		j := 2 * k // 在此轮循环中,data[k]和data[j]交换位置
		// 判断是否有右孩子        右孩子和左孩子比较大小
		if j+1 <= h.count && compare(h.data[h.indexes[j+1]], h.data[h.indexes[j]]) < 0 {
			j += 1 // 变更j为右孩子
		}
		if compare(h.data[h.indexes[k]], h.data[h.indexes[j]]) <= 0 {
			break // 满足二叉堆定义, 任意节点不小于其父节点
		}
		swapIndex(h.indexes, k, j) // 不满足二叉堆定义,就交换k和左右孩子中较小的孩子
		h.reverse[h.indexes[k]] = k
		h.reverse[h.indexes[j]] = j
		k = j // 继续进行下浮操作,直到满足二叉堆定义
	}
}

// 构造一个空堆,可容纳capacity个元素
func NewIndexMinHeap(capacity int) *IndexMinHeap {
	data := make([]interface{}, capacity+1)
	indexes := make([]int, capacity+1)
	reverse := make([]int, capacity+1)
	for i := 0; i <= capacity; i++ {
		reverse[i] = 0
	}
	return &IndexMinHeap{
		data:     data,
		indexes:  indexes,
		reverse:  reverse,
		count:    0,
		capacity: capacity,
	}
}

// 返回堆中的元素个数
func (h *IndexMinHeap) Size() int {
	return h.count
}

// 返回一个布尔值,表示堆中是否为空
func (h *IndexMinHeap) IsEmpty() bool {
	return h.count == 0
}

// 向最小堆中插入一个新的元素 item和元素的索引 i
// 传入的i对用户而言,是从0索引的
func (h *IndexMinHeap) Insert(item interface{}, i int) {
	if h.count+1 > h.capacity || i+1 < 1 || i+1 > h.capacity {
		panic("out of range.")
	}
	// 再插入一个新元素前,还需要保证索引i所在的位置是没有元素的
	if h.contain(i) {
		panic("该索引已经有元素!")
	}
	i += 1
	h.data[i] = item
	h.indexes[h.count+1] = i
	h.reverse[i] = h.count + 1
	h.count++

	h.shiftUp(h.count)
}

// 从最小索引堆中取出堆顶元素,即堆中所存储的最小数据
func (h *IndexMinHeap) ExtractMin() interface{} {
	if h.count <= 0 {
		panic("堆中没有元素!")
	}
	ret := h.data[h.indexes[1]]
	swapIndex(h.indexes, 1, h.count) // 最后一个元素和第一个元素交换
	h.reverse[h.indexes[h.count]] = 0
	h.count--
	if h.count > 0 {
		h.reverse[h.indexes[1]] = 1
		h.shiftDown(1)
	}
	return ret
}

// 从最小索引堆中取出堆顶元素的索引
func (h *IndexMinHeap) ExtractMinIndex() int {
	if h.count <= 0 {
		panic("堆中没有元素!")
	}
	ret := h.indexes[1] - 1
	swapIndex(h.indexes, 1, h.count)
	h.reverse[h.indexes[h.count]] = 0
	h.count--
	if h.count > 0 {
		h.reverse[h.indexes[1]] = 1
		h.shiftDown(1)
	}
	h.shiftDown(1)
	return ret
}

// 获取最小索引堆中的堆顶元素
func (h *IndexMinHeap) GetMin() interface{} {
	if h.count <= 0 {
		panic("堆中没有元素!")
	}
	return h.data[h.indexes[1]]
}

// 获取最小索引堆中堆顶元素的索引
func (h *IndexMinHeap) GetMinIndex() int {
	if h.count <= 0 {
		panic("堆中没有元素!")
	}
	return h.indexes[1] - 1
}

// 获取最小索引堆中索引为i的元素
func (h *IndexMinHeap) GetItem(i int) interface{} {
	if !h.contain(i) {
		panic("该索引没有元素!")
	}
	return h.data[i+1]
}

// 将最小索引堆中索引为i的元素修改为newItem
func (h *IndexMinHeap) contain(i int) bool {
	if i+1 < 1 || i+1 > h.capacity {
		panic("索引越界!")
	}
	return h.reverse[i+1] != 0
}

// 将最小索引堆中索引为i的元素修改为newItem
func (h *IndexMinHeap) change(i int, newItem interface{}) {
	if !h.contain(i) {
		panic("该索引没有元素!")
	}
	i += 1
	h.data[i] = newItem

	j := h.reverse[i]
	h.shiftUp(j)
	h.shiftDown(j)
}

// 交换堆中索引为i和j的两个元素
func swapIndex(data []int, i, j int) {
	data[i], data[j] = data[j], data[i]
}
