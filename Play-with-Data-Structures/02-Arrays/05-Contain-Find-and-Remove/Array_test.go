package _5_Contain_Find_and_Remove

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	arr := NewArray(20)
	for i := 0; i < 10; i++ {
		arr.AddLast(i)
	}
	fmt.Println(arr.ToString())

	arr.Add(1, 100)
	fmt.Println(arr.ToString())

	arr.AddFirst(-1)
	fmt.Println(arr.ToString())

	arr.Remove(2)
	fmt.Println(arr.ToString())

	arr.RemoveElement(4)
	fmt.Println(arr)

	arr.RemoveFirst()
	fmt.Println(arr)
}
