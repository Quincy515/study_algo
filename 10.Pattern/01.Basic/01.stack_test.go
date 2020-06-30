package main

import (
	"fmt"
	"testing"
)

// go 通过切片模拟栈
func TestStack(t *testing.T) {
	// 创建栈
	stack := make([]int, 0)
	// push 压入栈
	stack = append(stack, 10)
	// pop 弹出
	v := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	fmt.Println(v)
	// 检查栈空
	if len(stack) == 0 {
		return
	}
}
