package main

import (
	"testing"
)

var l *LinkedList

func init() {
	n5 := &ListNode{value: 5}
	n4 := &ListNode{value: 4, next: n5}
	n3 := &ListNode{value: 3, next: n4}
	n2 := &ListNode{value: 2, next: n3}
	n1 := &ListNode{value: 1, next: n2}
	l = &LinkedList{head: &ListNode{next: n1}}
}

func TestLinkedList_Reverse(t *testing.T) {
	l.Print() // 1->2->3->4->5
	l.Reverse()
	l.Print() // 5->4->3->2->1
}

func TestLinkedList_HasCycle(t *testing.T) {
	t.Log(l.HasCycle()) // false
	l.head.next.next.next.next.next.next = l.head.next.next.next
	t.Log(l.HasCycle()) // true
}

func TestMergeSortedList(t *testing.T) {
	n5 := &ListNode{value: 9}
	n4 := &ListNode{value: 7, next: n5}
	n3 := &ListNode{value: 5, next: n4}
	n2 := &ListNode{value: 3, next: n3}
	n1 := &ListNode{value: 1, next: n2}
	l1 := &LinkedList{head: &ListNode{next: n1}}

	n10 := &ListNode{value: 10}
	n9 := &ListNode{value: 8, next: n10}
	n8 := &ListNode{value: 6, next: n9}
	n7 := &ListNode{value: 4, next: n8}
	n6 := &ListNode{value: 2, next: n7}
	l2 := &LinkedList{head: &ListNode{next: n6}}

	MergeSortedList(l1, l2).Print() // 1->2->3->4->5->6->7->8->9->10
}

func TestLinkedList_DeleteBottomN(t *testing.T) {
	l.Print() // 1->2->3->4->5
	l.DeleteBottomN(3)
	l.Print() // 1->2->4->5
}

func TestLinkedList_FindMiddleNode(t *testing.T) {
	l.DeleteBottomN(1)
	l.DeleteBottomN(1)
	l.DeleteBottomN(1)
	l.DeleteBottomN(1)
	l.Print()
	t.Log(l.FindMiddleNode())
}
