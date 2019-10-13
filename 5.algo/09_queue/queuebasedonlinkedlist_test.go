package main

import "testing"

func TestLinkedListQueue_EnQueue(t *testing.T) {
	q := NewLinkedListQueue()
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)
	q.EnQueue(5)
	q.EnQueue(6)
	t.Log(q) // head<-<-1<-2<-3<-4<-5<-6<-tail
}

func TestLinkedListQueue_DeQueue(t *testing.T) {
	q := NewLinkedListQueue()
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)
	q.EnQueue(5)
	q.EnQueue(6)
	t.Log(q) // head<-<-1<-2<-3<-4<-5<-6<-tail
	q.DeQueue()
	t.Log(q) // head<-<-2<-3<-4<-5<-6<-tail
	q.DeQueue()
	t.Log(q) // head<-<-3<-4<-5<-6<-tail
	q.DeQueue()
	t.Log(q) // head<-<-4<-5<-6<-tail
	q.DeQueue()
	t.Log(q) // head<-<-5<-6<-tail
	q.DeQueue()
	t.Log(q) // head<-<-6<-tail
	q.DeQueue()
	t.Log(q) // empty queue
}
