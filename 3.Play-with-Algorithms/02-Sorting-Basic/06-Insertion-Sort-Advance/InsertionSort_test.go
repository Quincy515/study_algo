package Insertion_Sort

import (
	Selection_Sort "github.com/custergo/study_algo/3.Play-with-Algorithms/02-Sorting-Basic/04-Selection-Sort-Detect-Performance"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	N := 20000
	//arr := GenerateRandomArray(N, 0, N)
	arr := GenerateNearlyOrderArray(N, 100)
	arr2 := make([]int, N)
	copy(arr2, arr)
	TestSort("InsertionSort", InsertionSort, arr)
	TestSort("SelectionSort", Selection_Sort.SelectionSort, arr)
	arr = nil
	arr2 = nil
}
