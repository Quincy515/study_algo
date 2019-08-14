package Merge_Sort_Bottom_Up

import (
	Sort "github.com/custergo/study_algo/3.Play-with-Algorithms/02-Sorting-Basic/Chapter-02-Completed-Code"
	"testing"
)

func TestMergeSort(t *testing.T) {
	n := 50000
	//arr := GenerateRandomArray(n, 0, n)
	arr := GenerateNearlyOrderArray(n, 100)
	arr2 := make([]int, n)
	copy(arr2, arr)
	TestSort("InsertionSort", Sort.InsertionSort, arr)
	TestSort("MergeSortBU", MergeSortBU, arr2)
}
