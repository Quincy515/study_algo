package _4_LinkedList_and_Recursion

import (
	"fmt"
	"testing"
)

func TestListNode(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6}
	head := NewListNode(nums)
	fmt.Println(head)

	res4 := removeElements4(head, 6)
	fmt.Println(res4)

	res5 := removeElements5(head, 6)
	fmt.Println(res5)
}
