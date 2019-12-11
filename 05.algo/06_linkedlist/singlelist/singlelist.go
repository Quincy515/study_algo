package main

import "fmt"

// 单链表的基本操作

// 结点
type ListNode struct {
	next  *ListNode
	value interface{}
}

// 链表
type LinkedList struct {
	head   *ListNode // 头结点用来记录链表的基地址
	length uint
}

// 新建结点
func NewListNode(v interface{}) *ListNode {
	return &ListNode{nil, v}
}

// 后继结点 next 方法
func (this *ListNode) GetNext() *ListNode {
	return this.next // 指针的next 方法
}

// 获取当前结点值的方法
func (this *ListNode) GetValue() interface{} {
	return this.value
}

// 新建链表
func NewLinkedList() *LinkedList {
	return &LinkedList{NewListNode(0), 0}
}

// 在某个结点后面插入结点
func (this *LinkedList) InsertAfter(p *ListNode, v interface{}) bool {
	if nil == p {
		return false
	}
	newNode := NewListNode(v) // 新建一个结点即需要插入的结点
	oldNext := p.next         // 在p结点后面插入新结点，设置个临时变量存储
	p.next = newNode          // 在p结点后面插入新结点
	newNode.next = oldNext    // 新结点的后继结点指向之前p结点指向的结点
	this.length++             // 链表的长度增加
	return true               // 插入结点成功
}

// 在某个结点前面插入结点
func (this *LinkedList) InsertBefore(p *ListNode, v interface{}) bool {
	if nil == p || p == this.head {
		return false
	}
	cur := this.head.next
	pre := this.head
	for nil != cur {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}
	if nil == cur {
		return false
	}
	newNode := NewListNode(v)
	pre.next = newNode
	newNode.next = cur
	this.length++
	return true
}

// 在链表头部插入结点
func (this *LinkedList) InsertToHead(v interface{}) bool {
	return this.InsertAfter(this.head, v)
}

// 在链表尾部插入结点
func (this *LinkedList) InsertToTail(v interface{}) bool {
	cur := this.head
	for nil != cur.next {
		cur = cur.next
	}
	return this.InsertAfter(cur, v)
}

// 通过索引查找结点
func (this *LinkedList) FindByIndex(index uint) *ListNode {
	if index >= this.length {
		return nil
	}
	cur := this.head.next
	var i uint = 0
	for ; i < index; i++ {
		cur = cur.next
	}
	return cur
}

// 删除传入的结点
func (this *LinkedList) DeleteNode(p *ListNode) bool {
	if nil == p {
		return false
	}
	cur := this.head.next
	pre := this.head
	for nil != cur {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}
	if nil == cur {
		return false
	}
	pre.next = p.next
	p = nil
	this.length--
	return true
}

// 打印链表
func (this *LinkedList) Print() {
	cur := this.head.next
	format := ""
	for nil != cur {
		format += fmt.Sprintf("%+v", cur.GetValue())
		cur = cur.next
		if nil != cur {
			format += "->"
		}
	}
	fmt.Println(format)
}
