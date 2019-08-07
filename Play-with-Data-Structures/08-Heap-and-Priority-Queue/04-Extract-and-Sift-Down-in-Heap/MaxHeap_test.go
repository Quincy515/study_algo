package _4_Extract_and_Sift_Down_in_Heap

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestMaxHeap(t *testing.T) {
	n := 1000000 // 一百万

	maxHeap := NewMaxHeap(20)
	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		randNum := rand.Intn(math.MaxInt32)
		maxHeap.Add(randNum)
	}

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = maxHeap.ExtractMax().(int)
	}

	for i := 1; i < n; i++ {
		if arr[i-1] < arr[i] {
			panic("Error")
		}
	}

	fmt.Println("Test MaxHeap completed.")
}

/**
=== RUN   TestMaxHeap
Test MaxHeap completed.
--- PASS: TestMaxHeap (0.71s)
PASS
*/
