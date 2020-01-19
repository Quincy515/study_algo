package main

import (
	"fmt"
	"testing"
	"time"
)

func TestBinarySearch(t *testing.T) {
	n := 1000000
	data := generateOrderArray(n)

	startTime := time.Now()
	for i := 0; i < n; i++ {
		assert(i == binarySearch(data, n, i))
	}
	fmt.Println("binarySearch 7 complete.")
	fmt.Println("Time cost: ", time.Now().Sub(startTime))
}
