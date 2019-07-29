package _3_Add_Element_in_Array

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	arr := NewArray(5)
	fmt.Println(arr.GetSize())
	arr.AddLast(1)
	arr.AddLast(2)
	arr.AddLast(3)
	arr.Add(1, 5)
	fmt.Println(arr)
	fmt.Println(arr.GetSize())
}
