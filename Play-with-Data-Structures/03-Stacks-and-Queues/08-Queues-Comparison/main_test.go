package _8_Queues_Comparison

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func testQueue(queue Queue, opCount int) float64 {
	startTime := time.Now()
	for i := 0; i < opCount; i++ {
		queue.Enqueue(rand.Int())
	}
	for i := 0; i < opCount; i++ {
		queue.Dequeue()
	}
	return time.Now().Sub(startTime).Seconds()
}

func TestArrayLoopQueue(t *testing.T) {
	opCount := 100000

	arrayQueue := NewArrayQueue(20)
	time1 := testQueue(arrayQueue, opCount)
	fmt.Println("ArrayQueue, time:", time1, "s")

	loopQueue := NewLoopQueue(20)
	time2 := testQueue(loopQueue, opCount)
	fmt.Println("LoopQueue, time:", time2, "s")
}

/**
=== RUN   TestArrayLoopQueue
ArrayQueue, time: 8.244420569 s
LoopQueue, time: 0.01829418 s
--- PASS: TestArrayLoopQueue (8.26s)
PASS
*/
