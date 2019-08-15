package Max_Heap

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestMaxHeap(t *testing.T) {
	maxHeap := NewMaxHeap(100)
	//fmt.Println(maxHeap.Size())

	for i := 0; i < 100; i++ {
		maxHeap.Insert(rand.Int() % 100) // [0,100)随机数
	}
	//fmt.Println(maxHeap)

	for !maxHeap.IsEmpty() {
		fmt.Print(maxHeap.ExtractMax(), " ")
	}
}
