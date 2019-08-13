package Leetcode_347

type Freq struct {
	e, freq int
}

func (f *Freq) compareTo(another *Freq) int {
	if f.freq > another.freq {
		return -1
	} else if f.freq < another.freq {
		return 1
	} else {
		return 0
	}
}

type MaxHeap struct {
	data *Array
}

func NewMaxHeap(capacity int) *MaxHeap {
	return &MaxHeap{NewArray(capacity)}
}

// 数组转换成堆-heapify
func NewMaxHeapFromArr(arr []interface{}) *MaxHeap {
	maxHeap := &MaxHeap{NewArrayFromArr(arr)}
	for i := parent(len(arr) - 1); i >= 0; i-- {
		maxHeap.siftDown(i)
	}
	return maxHeap
}

// 返回堆中的元素个数
func (mh *MaxHeap) Size() int {
	return mh.data.GetSize()
}

// 返回一个布尔值，表示堆中是否为空
func (mh *MaxHeap) IsEmpty() bool {
	return mh.data.IsEmpty()
}

// 返回完全二叉树的数组表示中，一个索引所表示的元素的父亲节点的索引
func parent(index int) int {
	if index == 0 {
		panic("index-0 doesn't have parent.")
	}
	return (index - 1) / 2
}

// 返回完全二叉树的数组表示中，一个索引所表示的元素的左孩子节点的索引
func leftChild(index int) int {
	return index*2 + 1
}

// 返回完全二叉树的数组表示中，一个索引所表示的元素的右孩子节点的索引
func rightChild(index int) int {
	return index*2 + 2
}

// 向堆中添加元素
func (mh *MaxHeap) Add(e interface{}) {
	mh.data.AddLast(e)
	mh.siftUp(mh.data.GetSize() - 1)
}

func (mh *MaxHeap) siftUp(k int) {
	for k > 0 && mh.data.Get(parent(k)).(*Freq).compareTo(mh.data.Get(k).(*Freq)) < 0 {
		mh.data.Swap(k, parent(k))
		k = parent(k)
	}
}

// 看堆中的最大元素
func (mh *MaxHeap) FindMax() interface{} {
	if mh.data.GetSize() == 0 {
		panic("Can not findMax when heap is empty.")
	}
	return mh.data.Get(0) // 数组中索引为0的值
}

// 取出堆中最大元素
func (mh *MaxHeap) ExtractMax() interface{} {
	ret := mh.FindMax()

	mh.data.Swap(0, mh.data.GetSize()-1) // 数组索引为0和末尾元素交换值
	mh.data.RemoveLast()                 // 删除最后一个元素
	mh.siftDown(0)                       //需要满足堆的特性
	return ret
}

func (mh *MaxHeap) siftDown(k int) {
	for leftChild(k) < mh.data.GetSize() { // k的左孩子索引越界
		j := leftChild(k)
		// j+1是右孩子索引，如果存在右孩子比较后获得左右孩子中较大值的索引
		if j+1 < mh.data.GetSize() && mh.data.Get(j+1).(*Freq).compareTo(mh.data.Get(j).(*Freq)) > 0 {
			j++ // j 存储的就是右孩子的索引 j=rightChild(k)
		}
		// data[j] 是 leftChild 和 rightChild 中的最大值
		if mh.data.Get(k).(*Freq).compareTo(mh.data.Get(j).(*Freq)) > 0 {
			break // 下沉操作结束:满足堆中某个节点的值总是不大于其父亲节点的值
		}
		mh.data.Swap(k, j) // 不满足特性，交换位置
		k = j              // 进入下一轮循环，看是否进行下一轮下沉
	}
}

// 取出堆中的最大元素，并且替换成元素e
func (mh *MaxHeap) Replace(e interface{}) interface{} {
	ret := mh.FindMax()

	mh.data.Set(0, e)
	mh.siftDown(0)

	return ret
}
