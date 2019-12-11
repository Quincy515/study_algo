package Optimized_Selection_Sort

import (
	Selection_Sort "github.com/custergo/study_algo/03.Play-with-Algorithms/02-Sorting-Basic/04-Selection-Sort-Detect-Performance"
	Insertion_Sort "github.com/custergo/study_algo/03.Play-with-Algorithms/02-Sorting-Basic/06-Insertion-Sort-Advance"
	"testing"
)

func TestOptimizedSelectionSort(t *testing.T) {
	n := 20000
	arr := GenerateRandomArray(n, 0, n)
	//arr := GenerateNearlyOrderArray(N, 100)
	arr2 := make([]int, n)
	arr3 := make([]int, n)
	copy(arr2, arr)
	copy(arr3, arr)
	TestSort("InsertionSort", Insertion_Sort.InsertionSort, arr)
	TestSort("SelectionSort", Selection_Sort.SelectionSort, arr2)
	TestSort("OptimizedSelectionSort", OptimizedSelectionSort, arr3)
}
