package Sort

import (
	Optimized_Selection_Sort "github.com/custergo/study_algo/3.Play-with-Algorithms/02-Sorting-Basic/Optional-01-Optimized-Selection-Sort"
	"testing"
)

func TestOptimizedSelectionSort(t *testing.T) {
	n := 20000
	arr := GenerateRandomArray(n, 0, n)
	//arr := GenerateNearlyOrderArray(n, 100)
	arr2 := make([]int, n)
	arr3 := make([]int, n)
	arr4 := make([]int, n)
	arr5 := make([]int, n)
	copy(arr2, arr)
	copy(arr3, arr)
	copy(arr4, arr)
	copy(arr5, arr)
	TestSort("InsertionSort", InsertionSort, arr)
	TestSort("SelectionSort", SelectionSort, arr2)
	TestSort("OptimizedSelectionSort", Optimized_Selection_Sort.OptimizedSelectionSort, arr3)
	TestSort("BubbleSort", BubbleSort, arr4)
	TestSort("ShellSort", ShellSort, arr5)
}
