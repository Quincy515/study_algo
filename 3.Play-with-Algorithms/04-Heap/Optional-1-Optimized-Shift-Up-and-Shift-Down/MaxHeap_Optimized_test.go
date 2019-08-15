package Optimized_MaxHeap

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestIndexMaxHeap(t *testing.T) {
	maxIndexHeap := NewMaxHeap(100)

	for i := 0; i < 100; i++ {
		maxIndexHeap.Insert(rand.Int() % 100) // [0,100)随机数
	}
	//fmt.Println(maxIndexHeap)

	for !maxIndexHeap.IsEmpty() {
		fmt.Print(maxIndexHeap.ExtractMax(), " ")
	}
	fmt.Println()
}
