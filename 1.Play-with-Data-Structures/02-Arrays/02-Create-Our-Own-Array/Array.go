package _2_Create_Our_Own_Array

type Array struct {
	data []int
	size int
}

// 构造函数，传入数组的容量capacity构造Array
func NewArray(capacity int) *Array {
	// size默认为0
	return &Array{
		data: make([]int, capacity),
		size: 0,
	}
}

// 无参数的构造函数，默认数组的容量capacity=10
func NewArrayNP() *Array {
	return NewArray(10)
}

// 获取数组中的元素个数
func (a *Array) GetSize() int {
	return a.size
}

// 获取素组的容量
func (a *Array) GetCapacity() int {
	return len(a.data)
}

// 返回数组是否为空
func (a *Array) IsEmpty() bool {
	return a.size == 0
}
