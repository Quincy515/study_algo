package _7_Implementation_of_Loop_Queue

import (
	"fmt"
	"testing"
)

func TestLoopQueue(t *testing.T) {
	queue := NewLoopQueue(10)
	for i := 0; i < 10; i++ {
		queue.Enqueue(i)
		fmt.Println(queue)

		if i%3 == 2 {
			queue.Dequeue()
			fmt.Println(queue)
		}
	}
}
