package _5_Array_QueueArray

import (
	"fmt"
	"testing"
)

func TestArrayQueue(t *testing.T) {
	queue := NewArrayQueue(10)
	for i := 0; i < 10; i++ {
		queue.Enqueue(i)
		fmt.Println(queue)

		if i%3 == 2 {
			queue.Dequeue()
			fmt.Println(queue)
		}
	}
}
