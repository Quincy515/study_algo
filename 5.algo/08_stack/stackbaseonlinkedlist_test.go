package main

import "testing"

func TestLinkedListStack_Push(t *testing.T) {
	s := NewLinkedListStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Print() // 3 2 1
}

func TestLinkedListStack_Pop(t *testing.T) {
	s := NewLinkedListStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Print() // 3 2 1

	t.Log(s.Pop()) // 3
	t.Log(s.Pop()) // 2
	t.Log(s.Pop()) // 1
	t.Log(s.Pop()) // nil
	s.Print()      // empty stack
}

func TestLinkedListStack_Top(t *testing.T) {
	s := NewLinkedListStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	t.Log(s.Top()) // 3
	s.Pop()
	t.Log(s.Top()) // 2
	s.Pop()
	t.Log(s.Top()) // 1
	s.Pop()
	t.Log(s.Top()) // nil
}
