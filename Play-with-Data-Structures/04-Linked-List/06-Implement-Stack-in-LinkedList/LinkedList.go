package _6_Implement_Stack_in_LinkedList

import (
	"bytes"
	"fmt"
)

type Node struct {
	e    interface{}
	next *Node
}

func (n *Node) String() string {
	return fmt.Sprint(n.e)
}

type LinkedList struct {
	dummyHead *Node
	size      int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{dummyHead: &Node{}}
}

// 获取链表中的元素个数
func (l *LinkedList) GetSize() int {
	return l.size
}

// 返回链表是否为空
func (l *LinkedList) IsEmpty() bool {
	return l.size == 0
}

// 在链表头添加新的元素e
func (l *LinkedList) AddFirst(e interface{}) {
	//node := &Node{e, nil}
	//node.next = l.head
	//l.head = node

	//l.head = &Node{e, l.head}
	//l.size++
	l.Add(0, e)
}

// 在链表的index(0-based)位置添加新的元素e
// 在链表中不是一个常用的操作，学习练习使用
func (l *LinkedList) Add(index int, e interface{}) {
	if index < 0 || index > l.size {
		panic("Add failed. Illegal index.")
	}

	// 待插入节点的前一节点,从dummyHead开始
	prev := l.dummyHead
	for i := 0; i < index; i++ {
		prev = prev.next
	}

	// 插入新节点
	//node := &Node{e, nil}
	//node.next = prev.next
	//prev.next = node
	prev.next = &Node{e, prev.next}
	l.size++

}

// 在链表末尾添加新的元素e
func (l *LinkedList) AddLast(e interface{}) {
	l.Add(l.size, e)
}

// 获得链表的第index(0-bases)个位置的元素
// 在链表中不是一个常用的操作，练习用:)
func (l *LinkedList) Get(index int) interface{} {
	if index < 0 || index >= l.size {
		panic("Get failed. Illegal index.")
	}
	// 找index的元素从dummyHead.next开始
	cur := l.dummyHead.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	return cur.e
}

// 获得链表的第一个元素
func (l *LinkedList) GetFirst() interface{} {
	return l.Get(0)
}

// 获得链表的最后一个元素
func (l *LinkedList) GetLast() interface{} {
	return l.Get(l.size - 1)
}

// 修改链表的第index(o-based)个位置的元素为e
// 在链表中不是一个常用的操作，练习用:)
func (l *LinkedList) Set(index int, e interface{}) {
	if index < 0 || index >= l.size {
		panic("Set failed. Illegal index.")
	}
	cur := l.dummyHead.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	cur.e = e
}

// 查找链表中是否有元素e
func (l *LinkedList) Contains(e interface{}) bool {
	cur := l.dummyHead.next
	for cur != nil {
		if cur.e == e {
			return true
		}
		cur = cur.next
	}
	return false
}

// 从链表中删除index(0-based)位置的元素，返回删除的元素
// 在链表中不是一个常用的操作，练习用:)
func (l *LinkedList) Remove(index int) interface{} {
	if index < 0 || index >= l.size {
		panic("Remove failed. Illegal index.")
	}
	prev := l.dummyHead
	for i := 0; i < index; i++ {
		prev = prev.next
	}
	retNode := prev.next
	prev.next = retNode.next
	retNode.next = nil
	l.size--

	return retNode.e
}

// 从链表中删除第一个元素，返回删除的元素
func (l *LinkedList) RemoveFirst() interface{} {
	return l.Remove(0)
}

// 从链表中删除最后一个元素，返回删除的元素
func (l *LinkedList) RemoveLast() interface{} {
	return l.Remove(l.size - 1)
}

func (l *LinkedList) String() string {
	buffer := bytes.Buffer{}
	//cur := l.dummyHead.next
	//for cur != nil {
	//	buffer.WriteString(fmt.Sprint(cur.e) + "->")
	//	cur = cur.next
	//}
	for curr := l.dummyHead.next; curr != nil; curr = curr.next {
		buffer.WriteString(fmt.Sprint(curr.e) + "->")
	}
	buffer.WriteString("NULL")
	return buffer.String()
}
