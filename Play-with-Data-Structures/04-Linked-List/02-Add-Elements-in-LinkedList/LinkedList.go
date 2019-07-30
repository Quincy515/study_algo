package _2_Add_Elements_in_LinkedList

import "fmt"

type Node struct {
	e    interface{}
	next *Node
}

func (n *Node) String() string {
	return fmt.Sprint(n.e)
}

type LinkedList struct {
	head *Node
	size int
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

	l.head = &Node{e, l.head}
	l.size++
}

// 在链表的index(0-based)位置添加新的元素e
// 在链表中不是一个常用的操作，学习练习使用
func (l *LinkedList) Add(index int, e interface{}) {
	if index < 0 || index > l.size {
		panic("Add failed. Illegal index.")
	}
	if index == 0 {
		l.AddFirst(e)
	} else {
		// 待插入节点的前一节点
		prev := l.head
		for i := 0; i < index-1; i++ {
			prev = prev.next
		}

		// 插入新节点
		//node := &Node{e, nil}
		//node.next = prev.next
		//prev.next = node
		prev.next = &Node{e, prev.next}
		l.size++
	}
}

// 在链表末尾添加新的元素e
func (l *LinkedList) AddLast(e interface{}) {
	l.Add(l.size, e)
}
