package Lazy_Prim

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestIndexMaxHeap(t *testing.T) {
	maxIndexHeap := NewMinHeap(100)

	for i := 0; i < 100; i++ {
		maxIndexHeap.Insert(rand.Int() % 100) // [0,100)随机数
	}
	//fmt.Println(maxIndexHeap)

	for !maxIndexHeap.IsEmpty() {
		fmt.Print(maxIndexHeap.ExtractMin(), " ")
	}
	fmt.Println()
}
