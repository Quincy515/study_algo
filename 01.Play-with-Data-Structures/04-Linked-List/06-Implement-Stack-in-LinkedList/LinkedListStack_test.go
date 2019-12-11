package _6_Implement_Stack_in_LinkedList

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestLinkedListStack(t *testing.T) {
	stack := NewLinkedListStack()

	for i := 0; i < 5; i++ {
		stack.Push(i)
		fmt.Println(stack)
	}

	stack.Pop()
	fmt.Println(stack)
}

func testStack(stack Stack, opCount int) float64 {
	startTime := time.Now()
	for i := 0; i < opCount; i++ {
		stack.Push(rand.Int())
	}
	for i := 0; i < opCount; i++ {
		stack.Pop()
	}

	return time.Now().Sub(startTime).Seconds()
}

func TestLinkedListArrayStack(t *testing.T) {
	opCount := 10000000
	arrayStack := NewArrayStack(20)
	time1 := testStack(arrayStack, opCount)
	fmt.Println("ArrayStack time: ", time1, " s")

	linkedListStack := NewLinkedListStack()
	time2 := testStack(linkedListStack, opCount)
	fmt.Println("LinkedListStack, time: ", time2, " s")
}

/**
其实这个时间比较很复杂，因为 LinkedListStack 中包含更多的 new 操作
=== RUN   TestLinkedListArrayStack
ArrayStack time:  1.431879572  s
LinkedListStack, time:  1.755778205  s
--- PASS: TestLinkedListArrayStack (3.19s)
PASS
*/
