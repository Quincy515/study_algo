package _3_Add_Element_in_Array

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

// 向所有元素后添加一个新元素
func (a *Array) AddLast(e int) {
	//if a.size == len(a.data) {
	//	panic("AddList failed. Array is full.")
	//}
	//a.data[a.size] = e
	//a.size++
	a.Add(a.size, e)
}

// 在第index个位置插入一个新元素e
func (a *Array) Add(index, e int) {
	if a.size == len(a.data) {
		panic("AddList failed. Array is full.")
	}
	if index < 0 || index > a.size {
		panic("AddList failed. Require index >= 0 and index <= size.")
	}

	for i := a.size - 1; i >= index; i-- {
		a.data[i+1] = a.data[i]
	}
	a.data[index] = e
	a.size++
}
