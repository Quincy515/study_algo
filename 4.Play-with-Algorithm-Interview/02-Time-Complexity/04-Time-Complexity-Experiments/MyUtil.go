package main

import (
	"math/rand"
	"time"
)

// 下面是辅助函数
// 生成给定范围的无序数组
func generateRandomArray(n, rangeL, rangeR int) []int {
	assert(n > 0 && rangeL <= rangeR)

	arr := make([]int, n)

	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		arr[i] = rand.Int()%(rangeR-rangeL+1) + rangeL
	}
	return arr
}

// 生成从小到大有序的数组
func generateOrderArray(n int) []int {
	assert(n > 0)

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = i
	}
	return arr
}

// 断言函数
func assert(exp bool) {
	if !exp {
		panic("断言失败！")
	}
}

// 求最小值
func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// 交换数组中下标为i和j的值
func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
