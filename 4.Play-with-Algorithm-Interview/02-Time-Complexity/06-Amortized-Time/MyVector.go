package main

type MyVector struct {
	data     []interface{}
	capacity int // 存储数组中可以容纳的最大的元素个数
	size     int // 存储数组中的元素个数
}

func NewMyVector() *MyVector {
	return &MyVector{data: make([]interface{}, 100), capacity: 100, size: 0}
}

// 平均复杂度为 O(1)
func (this *MyVector) push_back(e interface{}) {
	if this.size == this.capacity {
		this.resize(2 * this.capacity)
	}
	this.data[this.size] = e
	this.size++
}

// 复杂度为O(n)
func (this *MyVector) resize(newCapacity int) {
	if newCapacity < this.size {
		panic("newCapacity 不能小于 size")
	}
	newData := make([]interface{}, newCapacity)
	for i := 0; i < this.size; i++ {
		newData[i] = this.data[i]
	}
	this.data = newData
	this.capacity = newCapacity
}

// 平均复杂度为 O(1)
func (this MyVector) pop_back() interface{} {
	if this.size <= 0 {
		panic("vector不能为空!")
	}
	this.size--
	return this.data[this.size]
}
