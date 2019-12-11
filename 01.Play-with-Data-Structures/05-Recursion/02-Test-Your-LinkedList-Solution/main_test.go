package _2_Test_Your_LinkedList_Solution

import (
	"fmt"
	"testing"
)

func TestListNode(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6}
	head := NewListNode(nums)
	fmt.Println(head)

	res1 := removeElements1(head, 6)
	fmt.Println(res1)

	res2 := removeElements2(head, 6)
	fmt.Println(res2)

	res3 := removeElements3(head, 6)
	fmt.Println(res3)
}
