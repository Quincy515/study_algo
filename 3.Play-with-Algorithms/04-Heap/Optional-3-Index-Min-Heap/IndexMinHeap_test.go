package Index_Min_Heap

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestIndexMaxHeap(t *testing.T) {
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
