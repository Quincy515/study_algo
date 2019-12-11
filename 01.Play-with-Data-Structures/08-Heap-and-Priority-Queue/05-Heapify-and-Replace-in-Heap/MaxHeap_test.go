package _5_Heapify_and_Replace_in_Heap

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
)

func testHeap(testData []interface{}, isHeapify bool) time.Duration {
	startTime := time.Now()

	var maxHeap *MaxHeap
	if isHeapify {
		maxHeap = NewMaxHeapFromArr(testData)
	} else {
		maxHeap = NewMaxHeap(20)
		for _, v := range testData {
			maxHeap.Add(v)
		}
	}

	arr := make([]interface{}, len(testData))
	for i := 0; i < len(testData); i++ {
		arr[i] = maxHeap.ExtractMax()
	}
	for i := 1; i < len(testData); i++ {
		if compare(arr[i-1], arr[i]) < 0 {
			panic("Error")
		}
	}
	fmt.Println("Test MaxHeap completed.")

	return time.Now().Sub(startTime)
}
func TestMaxHeap(t *testing.T) {
	n := 1000000 // 一百万

	var testData []interface{}
	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		testData = append(testData, rand.Intn(math.MaxInt32))
	}

	time1 := testHeap(testData, false)
	fmt.Println("Without heapify: ", time1)

	time2 := testHeap(testData, true)
	fmt.Println("With heapify: ", time2)
}

/**
=== RUN   TestMaxHeap
Test MaxHeap completed.
Without heapify:  953.559956ms   O(nlogn)
Test MaxHeap completed.
With heapify:  866.746395ms      O(n)
--- PASS: TestMaxHeap (2.06s)
PASS
*/
