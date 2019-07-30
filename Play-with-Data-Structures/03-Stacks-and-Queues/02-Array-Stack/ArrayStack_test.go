package _2_Array_Stack

import (
	"fmt"
	"testing"
)

func TestArrayStack(t *testing.T) {
	stack := NewArrayStackNo()

	for i := 0; i < 5; i++ {
		stack.Push(i)
		fmt.Println(stack)
	}

	stack.Pop()
	fmt.Println(stack)
}
