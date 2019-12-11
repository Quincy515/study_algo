package main

import (
	"math/rand"
	"reflect"
	"time"
)

// 下面是辅助函数
// 生成给定范围的无序数组
func generateRandomArray(n, rangeL, rangeR int) []interface{} {
	assert(n > 0 && rangeL <= rangeR)

	arr := make([]interface{}, n)

	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		arr[i] = rand.Int()%(rangeR-rangeL+1) + rangeL
	}
	return arr
}

// 生成从小到大有序的数组
func generateOrderArray(n int) []interface{} {
	assert(n > 0)

	arr := make([]interface{}, n)

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
func min(a, b interface{}) interface{} {
	if compare(a, b) <= 0 {
		return a
	}
	return b
}

// 交换数组中下标为i和j的值
func swap(arr []interface{}, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// 比较两个值的大小
func compare(a, b interface{}) int {
	aType := reflect.TypeOf(a).String()
	bType := reflect.TypeOf(b).String()
	if aType != bType {
		panic("cannot compare different type params")
	}

	switch a.(type) {
	case int:
		if a.(int) > b.(int) {
			return 1
		} else if a.(int) < b.(int) {
			return -1
		} else {
			return 0
		}
	case float64:
		if a.(float64) > b.(float64) {
			return 1
		} else if a.(float64) < b.(float64) {
			return -1
		} else {
			return 0
		}
	case string:
		if a.(string) > b.(string) {
			return 1
		} else if a.(string) < b.(string) {
			return -1
		} else {
			return 0
		}
	default:
		panic("unsupported type params")
	}
}
