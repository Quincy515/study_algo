package _4_Query_and_Update_Element

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
}
