package _4_LinkedList_and_Recursion

import (
	"bytes"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 链表节点的构造函数
// 使用arr为参数，创建一个链表，当前的ListNode
func NewListNode(arr []int) *ListNode {
	if arr == nil || len(arr) == 0 {
		panic("arr can not be empty.")
	}

	head := &ListNode{Val: arr[0]} // 创建头节点
	cur := head
	for i := 1; i < len(arr); i++ {
		cur.Next = &ListNode{Val: arr[i]}
		cur = cur.Next
	}

	return head
}

// 以当前节点为头节点的链表信息字符串
func (l *ListNode) String() string {
	var buffer bytes.Buffer
	for l != nil {
		buffer.WriteString(fmt.Sprintf("%v->", l.Val))
		l = l.Next
	}
	buffer.WriteString("NULL")
	return buffer.String()
}
