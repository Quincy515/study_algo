package main

import "testing"

func TestArrayQueue_EnQueue(t *testing.T) {
	q := NewArrayQueue(5)
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)
	q.EnQueue(5)
	q.EnQueue(6)
	t.Log(q) // head<-1<-2<-3<-4<-5<-tail
}

func TestArrayQueue_DeQueue(t *testing.T) {
	q := NewArrayQueue(5)
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)
	q.EnQueue(5)
	q.EnQueue(6)
	t.Log(q) // head<-1<-2<-3<-4<-5<-tail
	q.DeQueue()
	t.Log(q) // head<-2<-3<-4<-5<-tail
	q.DeQueue()
	t.Log(q) // head<-3<-4<-5<-tail
	q.DeQueue()
	t.Log(q) // head<-4<-5<-tail
	q.DeQueue()
	t.Log(q) // head<-5<-tail
	q.DeQueue()
	t.Log(q) // empty queue
}
