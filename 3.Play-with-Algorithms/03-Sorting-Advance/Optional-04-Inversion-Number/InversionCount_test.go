package Inversion_Number

import (
	"fmt"
	"testing"
)

func TestInversionCount(t *testing.T) {
	n := 1000000

	// 测试1: 测试随机数组
	arr := GenerateRandomArray(n, 0, n)
	fmt.Println("Test Inversion Count for Random Array: ", InversionCount(arr, n))

	// 测试2: 测试完全有序的数组
	// 结果应为0
	arr = GenerateOrderedArray(n)
	fmt.Println("Test Inversion Count for Ordered Array: ", InversionCount(arr, n))

	// 测试3: 测试完全能逆序的数组
	// 结果应该为 N*(N-1)/2 = 499,999,500,000
	arr = GenerateInversedArray(n)
	fmt.Println("Test Inversion Count for Ordered Array: ", InversionCount(arr, n))

}
