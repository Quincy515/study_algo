package _6_Debug_Recursive_Solution

import (
	"bytes"
	"fmt"
)

func removeElements(head *ListNode, val, depth int) *ListNode {
	// depthString 代表递归深度
	depthString := generateDepthString(depth)
	fmt.Print(depthString) // 递归调用前打印
	fmt.Println("Call: remove ", val, " in ", head)

	if head == nil {
		fmt.Print(depthString) // 递归调用前打印
		fmt.Println("Return: ", head)
		return nil
	}
	res := removeElements(head.Next, val, depth+1)
	fmt.Print(depthString) // 递归调用后打印
	fmt.Println("After remove ", val, ": ", res)

	ret := &ListNode{}
	if head.Val == val {
		ret = res
	} else {
		head.Next = res
		ret = head
	}
	fmt.Print(depthString) // 节点处理后打印
	fmt.Println("Return: ", ret)

	return ret
}

func generateDepthString(depth int) string {
	var buffer bytes.Buffer
	for i := 0; i < depth; i++ {
		buffer.WriteString("--")
	}
	return buffer.String()
}
