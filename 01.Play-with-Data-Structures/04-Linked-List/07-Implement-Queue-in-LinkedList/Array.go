package _7_Implement_Queue_in_LinkedList

import (
	"bytes"
	"fmt"
)

type Array struct {
	data []interface{}
	size int
}

// 构造函数，传入数组的容量capacity构造Array
func NewArray(capacity int) *Array {
	// size默认为0
	return &Array{
		data: make([]interface{}, capacity),
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
func (a *Array) AddLast(e interface{}) {
	a.Add(a.size, e)
}

// 向所有元素前添加一个新元素
func (a *Array) AddFirst(e interface{}) {
	a.Add(0, e)
}

// 在第index个位置插入一个新元素e
func (a *Array) Add(index int, e interface{}) {

	if index < 0 || index > a.size {
		panic("AddList failed. Require index >= 0 and index <= size.")
	}
	if a.size == len(a.data) {
		// 数组空间的扩容，现在数组空间的2倍
		a.resize(2 * len(a.data))
	}
	for i := a.size - 1; i >= index; i-- {
		a.data[i+1] = a.data[i]
	}
	a.data[index] = e
	a.size++
}

// 获取index索引位置的元素
func (a *Array) Get(index int) interface{} {
	if index < 0 || index >= a.size {
		panic("Get failed. Index is illegal.")
	}
	return a.data[index]
}

func (a *Array) GetLast() interface{} {
	return a.Get(a.size - 1)
}

func (a *Array) GetFirst() interface{} {
	return a.Get(0)
}

// 修改index索引位置的元素
func (a *Array) Set(index int, e interface{}) {
	if index < 0 || index >= a.size {
		panic("Get failed. Index is illegal.")
	}
	a.data[index] = e
}

// 查找数组中是否有元素e
func (a *Array) Contains(e interface{}) bool {
	for i := 0; i < a.size; i++ {
		if a.data[i] == e {
			return true
		}
	}
	return false
}

// 查找数组中元素e所在的索引，如果不存在元素e, 则返回-1
func (a *Array) Find(e interface{}) int {
	for i := 0; i < a.size; i++ {
		if a.data[i] == e {
			return i
		}
	}
	return -1
}

// 从数组中删除index位置的元素，返回删除的元素
func (a *Array) Remove(index int) interface{} {
	if index < 0 || index >= a.size {
		panic("Get failed. Index is illegal.")
	}

	ret := a.data[index]
	for i := index + 1; i < a.size; i++ {
		a.data[i-1] = a.data[i]
	}
	a.size--
	a.data[a.size] = nil // loitering object != memory leak

	if a.size == len(a.data)/4 && len(a.data)/2 != 0 {
		a.resize(len(a.data) / 2)
	}
	return ret
}

// 从数组中删除第一个元素，返回删除的元素
func (a *Array) RemoveFirst() interface{} {
	return a.Remove(0)
}

// 从数组中删除最后一个元素，返回删除的元素
func (a *Array) RemoveLast() interface{} {
	return a.Remove(a.size - 1)
}

// 从数组中删除元素e
func (a *Array) RemoveElement(e interface{}) bool {
	index := a.Find(e)
	if index != -1 {
		a.Remove(index)
		return true
	}
	return false
}

func (a *Array) ToString() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Array: size = %d, capacity = %d\n", a.size, len(a.data)))
	buffer.WriteString("[")
	for i := 0; i < a.size; i++ {
		// fmt.Sprint将interface{}类型转换为字符串
		buffer.WriteString(fmt.Sprintf("%v", a.data[i]))
		if i != (a.size - 1) {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("]")
	return buffer.String()
}

// 重写Array的string方法
func (a *Array) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Array: size = %d, capacity = %d\n", a.size, len(a.data)))
	buffer.WriteString("[")
	for i := 0; i < a.size; i++ {
		// fmt.Sprint将interface{}类型转换为字符串
		buffer.WriteString(fmt.Sprintf("%v", a.data[i]))
		if i != (a.size - 1) {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("]")
	return buffer.String()
}

func (a *Array) resize(newCapacity int) {
	newData := make([]interface{}, newCapacity)
	for i := 0; i < a.size; i++ {
		newData[i] = a.data[i]
	}
	a.data = newData
}
