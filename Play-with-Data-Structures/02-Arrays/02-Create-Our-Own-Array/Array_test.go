package _2_Create_Our_Own_Array

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	arr := NewArray(5)
	fmt.Println(arr)
	fmt.Println(arr.GetCapacity(), arr.GetSize(), arr.IsEmpty())
}
