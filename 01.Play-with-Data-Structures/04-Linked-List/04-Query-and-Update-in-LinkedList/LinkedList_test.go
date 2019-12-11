package _4_Query_and_Update_in_LinkedList

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
}
