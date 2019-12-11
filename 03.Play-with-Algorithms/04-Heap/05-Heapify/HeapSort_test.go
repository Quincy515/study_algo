package Heap_Sort

import (
	"fmt"
	Sort "github.com/custergo/study_algo/03.Play-with-Algorithms/03-Sorting-Advance/Optional-03-ShellSort-MergeSort-and-QuickSort-Comparision"
	"testing"
)

func TestHeapSort1(t *testing.T) {
	n := 1000000
	// 测试1: 一般性能测试
	fmt.Println("Test for Random Array, size= ", n, " random range [0, ", n, "]")
	arr1 := GenerateRandomArray(n, 0, n)
	arr2 := make([]int, n)
	arr3 := make([]int, n)
	arr4 := make([]int, n)
	arr5 := make([]int, n)
	copy(arr2, arr1)
	copy(arr3, arr1)
	copy(arr4, arr1)
	copy(arr5, arr1)

	TestSort("MergeSort", Sort.MergeSort, arr1)
	TestSort("QuickSort", Sort.QuickSort, arr2)
	TestSort("QuickSort3Ways", Sort.QuickSort3Ways, arr3)
	TestSort("HeapSort1", HeapSort1, arr4)
	TestSort("HeapSort2", HeapSort2, arr5)

	// 测试2: 测试近乎有序的数组
	swapTimes := 100
	fmt.Println("Test for Random Nearly Ordered Array, size= ", n, " swap time= ", swapTimes)
	arr1 = GenerateNearlyOrderArray(n, swapTimes)
	copy(arr2, arr1)
	copy(arr3, arr1)
	copy(arr4, arr1)
	copy(arr5, arr1)

	TestSort("MergeSort", Sort.MergeSort, arr1)
	TestSort("QuickSort", Sort.QuickSort, arr2)
	TestSort("QuickSort3Ways", Sort.QuickSort3Ways, arr3)
	TestSort("HeapSort1", HeapSort1, arr4)
	TestSort("HeapSort2", HeapSort2, arr5)

	// 测试3: 测试存在包含大量相同元素的数组
	fmt.Println("Test for Random Array, size= ", n, " random range [0, 10]")
	arr1 = GenerateRandomArray(n, 0, 10)
	copy(arr2, arr1)
	copy(arr3, arr1)
	copy(arr4, arr1)
	copy(arr5, arr1)

	TestSort("MergeSort", Sort.MergeSort, arr1)
	TestSort("QuickSort", Sort.QuickSort, arr2)
	TestSort("QuickSort3Ways", Sort.QuickSort3Ways, arr3)
	TestSort("HeapSort1", HeapSort1, arr4)
	TestSort("HeapSort2", HeapSort2, arr5)
}
