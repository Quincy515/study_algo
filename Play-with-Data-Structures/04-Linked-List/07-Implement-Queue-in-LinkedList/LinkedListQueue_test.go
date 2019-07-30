package _7_Implement_Queue_in_LinkedList

import (
	"fmt"
	"testing"
)

func TestLinkedListQueue(t *testing.T) {
	queue := NewLinkedListQueue()

	for i := 0; i < 10; i++ {
		queue.Enqueue(i)
		fmt.Println(queue)

		if i%3 == 2 {
			queue.Dequeue()
			fmt.Println(queue)
		}
	}
}
