package _5_Contain_Find_and_Remove

import (
	"bytes"
	"fmt"
	"strconv"
)

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

// 向所有元素前添加一个新元素
func (a *Array) AddFirst(e int) {
	a.Add(0, e)
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

// 获取index索引位置的元素
func (a *Array) Get(index int) int {
	if index < 0 || index >= a.size {
		panic("Get failed. Index is illegal.")
	}
	return a.data[index]
}

// 修改index索引位置的元素
func (a *Array) Set(index, e int) {
	if index < 0 || index >= a.size {
		panic("Get failed. Index is illegal.")
	}
	a.data[index] = e
}

// 查找数组中是否有元素e
func (a *Array) Contains(e int) bool {
	for i := 0; i < a.size; i++ {
		if a.data[i] == e {
			return true
		}
	}
	return false
}

// 查找数组中元素e所在的索引，如果不存在元素e, 则返回-1
func (a *Array) Find(e int) int {
	for i := 0; i < a.size; i++ {
		if a.data[i] == e {
			return i
		}
	}
	return -1
}

// 查找数组中元素e所在的索引组成的切片，不存在则返回-1
//func (a *Array) FindAll(e int) (index []int) {
//	for i := 0; i < a.size; i++ {
//		if Interfaces.Compare(a.data[i], e) == 0 {
//			index = append(index, i)
//		}
//	}
//	return
//}

// 从数组中删除index位置的元素，返回删除的元素
func (a *Array) Remove(index int) int {
	if index < 0 || index >= a.size {
		panic("Get failed. Index is illegal.")
	}

	ret := a.data[index]
	for i := index + 1; i < a.size; i++ {
		a.data[i-1] = a.data[i]
	}
	a.size--
	return ret
}

// 从数组中删除第一个元素，返回删除的元素
func (a *Array) RemoveFirst() int {
	return a.Remove(0)
}

// 从数组中删除最后一个元素，返回删除的元素
func (a *Array) RemoveLast() int {
	return a.Remove(a.size - 1)
}

// 从数组中删除元素e
func (a *Array) RemoveElement(e int) {
	index := a.Find(e)
	if index != -1 {
		a.Remove(index)
	}
}

// 从数组中删除所有元素e
//func (a *Array) RemoveAllElement(e int) bool {
//	index := a.find(e)
//	if index == -1 {
//		return false
//	}
//	for i := 0; i < a.size; i++ {
//		if Interfaces.Compare(a.data[i], e) == 0 {
//			a.Remove(i)
//		}
//	}
//	return true
//}

func (a *Array) ToString() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Array: size = %d, capacity = %d\n", a.size, len(a.data)))
	buffer.WriteString("[")
	for i := 0; i < a.size; i++ {
		buffer.WriteString(strconv.Itoa(a.data[i]))
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
		buffer.WriteString(strconv.Itoa(a.data[i]))
		if i != (a.size - 1) {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("]")
	return buffer.String()
}
