package main

import (
	"fmt"
	"math"
	"testing"
	"time"
)

func TestSort(t *testing.T) {
	// O(N) 数据规模倍乘测试 findMax
	fmt.Println("Test for findMax: ")
	for i := 10; i <= 26; i++ {
		n := int(math.Pow(float64(2), float64(i)))
		arr := generateRandomArray(n, 0, 100000000)

		startTime := time.Now()
		findMax(arr, n)
		fmt.Println("data size 2^", i, " = ", n) // 2倍的运行时间
		fmt.Println("Time cost: ", time.Now().Sub(startTime))
	}
	fmt.Println()

	// O(N^2) 数据规模倍乘测试selectionSort
	fmt.Println("Test for Selection Sort:")
	for i := 10; i <= 16; i++ {
		n := int(math.Pow(float64(2), float64(i)))
		arr := generateRandomArray(n, 0, 100000000)

		startTime := time.Now()
		selectionSort(arr, n)
		fmt.Println("data size 2^", i, " = ", n) // 4倍的运行时间
		fmt.Println("Time cost: ", time.Now().Sub(startTime))
	}
	fmt.Println()

	// O(logN) 数据规模倍乘测试binarySearch
	fmt.Println("Test for binarySearch:")
	for i := 10; i <= 26; i++ {
		n := int(math.Pow(float64(2), float64(i)))
		arr := generateOrderArray(n)

		startTime := time.Now()
		binarySearch(arr, n, 0)
		fmt.Println("data size 2^", i, " = ", n) // (1+log2/logN)倍的运行时间
		fmt.Println("Time cost: ", time.Now().Sub(startTime))
	}
	fmt.Println()

	// O(NlogN) 数据规模倍乘测试mergeSort
	fmt.Println("Test for Merge Sort:")
	for i := 10; i <= 26; i++ {
		n := int(math.Pow(float64(2), float64(i)))
		arr := generateRandomArray(n, 0, 1<<30)

		startTime := time.Now()
		mergeSort(arr, n)
		fmt.Println("data size 2^", i, " = ", n) // 2倍多一点的运行时间
		fmt.Println("Time cost: ", time.Now().Sub(startTime))
	}
	fmt.Println()
}
