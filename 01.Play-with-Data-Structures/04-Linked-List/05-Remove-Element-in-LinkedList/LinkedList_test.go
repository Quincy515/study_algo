package _5_Remove_Element_in_LinkedList

import (
	"fmt"
	"testing"
)

func TestLinkedList(t *testing.T) {
	linkedList := NewLinkedList()
	for i := 0; i < 5; i++ {
		linkedList.AddFirst(i)
		fmt.Println(linkedList)
	}

	linkedList.Add(2, 666)
	fmt.Println(linkedList)

	linkedList.Remove(2)
	fmt.Println(linkedList)

	linkedList.RemoveFirst()
	fmt.Println(linkedList)

	linkedList.RemoveLast()
	fmt.Println(linkedList)
}
