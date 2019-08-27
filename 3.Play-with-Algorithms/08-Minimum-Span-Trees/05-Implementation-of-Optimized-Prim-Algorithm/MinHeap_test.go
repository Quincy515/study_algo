package Prim

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

	minIndexHeap := NewIndexMinHeap(100)

	for i := 0; i < 100; i++ {
		minIndexHeap.Insert(rand.Int()%100, i) // [0,100)随机数
	}
	//fmt.Println(maxIndexHeap)

	for !minIndexHeap.IsEmpty() {
		fmt.Print(minIndexHeap.ExtractMin(), " ")
	}
	fmt.Println()
}
