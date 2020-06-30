package main

import (
	"fmt"
	"testing"
)

// go 通过切片模拟队列
func TestQueue(t *testing.T){
	// 创建队列
	queue := make([]int, 0)
	// enqueue 入队
	queue = append(queue, 10)
	// dequeue 出队
	v := queue[0]
	queue = queue[1:]
	fmt.Println(v)
	// 长度0为空
	if len(queue) == 0 {
		return
	}
}