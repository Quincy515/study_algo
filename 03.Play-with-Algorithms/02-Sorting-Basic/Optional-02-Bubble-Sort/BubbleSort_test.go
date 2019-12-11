package Bubble_Sort

import (
	Selection_Sort "github.com/custergo/study_algo/03.Play-with-Algorithms/02-Sorting-Basic/04-Selection-Sort-Detect-Performance"
	Insertion_Sort "github.com/custergo/study_algo/03.Play-with-Algorithms/02-Sorting-Basic/06-Insertion-Sort-Advance"
	Optimized_Selection_Sort "github.com/custergo/study_algo/03.Play-with-Algorithms/02-Sorting-Basic/Optional-01-Optimized-Selection-Sort"
	"testing"
)

func TestOptimizedSelectionSort(t *testing.T) {
	n := 20000
	arr := GenerateRandomArray(n, 0, n)
	//arr := GenerateNearlyOrderArray(N, 100)
	arr2 := make([]int, n)
	arr3 := make([]int, n)
	arr4 := make([]int, n)
	arr5 := make([]int, n)
	arr6 := make([]int, n)
	copy(arr2, arr)
	copy(arr3, arr)
	copy(arr4, arr)
	copy(arr5, arr)
	copy(arr6, arr)
	TestSort("InsertionSort", Insertion_Sort.InsertionSort, arr)
	TestSort("SelectionSort", Selection_Sort.SelectionSort, arr2)
	TestSort("OptimizedSelectionSort", Optimized_Selection_Sort.OptimizedSelectionSort, arr3)
	TestSort("BubbleSort1", BubbleSort1, arr4)
	TestSort("BubbleSort2", BubbleSort2, arr5)
	TestSort("BubbleSort3", BubbleSort3, arr6)
}
