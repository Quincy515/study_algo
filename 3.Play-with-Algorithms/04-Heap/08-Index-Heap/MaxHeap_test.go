package Index_Heap

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestIndexMaxHeap(t *testing.T) {
	maxIndexHeap := NewIndexMaxHeap(100)
	fmt.Println(maxIndexHeap.Size())

	for i := 0; i < 100; i++ {
		maxIndexHeap.Insert(rand.Int()%100, i) // [0,100)随机数
	}
	//fmt.Println(maxIndexHeap)

	for !maxIndexHeap.IsEmpty() {
		fmt.Print(maxIndexHeap.ExtractMax(), " ")
	}
}
