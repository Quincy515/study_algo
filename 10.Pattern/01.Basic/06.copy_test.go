package main

import "testing"

func TestCopy(t *testing.T) {
	// 删除 a[i]，可以用 copy 将 i+1 到末尾的值覆盖到 i
	a := []int{1, 2, 3, 4}
	i := 0
	copy(a[i:], a[i+1:])
	t.Log(a) // 2, 3, 4, 4

	// make 创建长度，通过索引赋值
	arr := make([]int, len(a))
	arr[len(a)-1] = a[len(a)-1]
	t.Log(arr) // 0, 0, 0, 4

	// make 长度为0，则通过 append() 赋值
	arr = make([]int, 0)
	arr = append(arr, a[len(a)-1])
	t.Log(arr) // 4
}
