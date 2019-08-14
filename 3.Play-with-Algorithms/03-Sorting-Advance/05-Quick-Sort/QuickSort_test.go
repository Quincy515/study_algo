package Quick_Sort

import (
	Merge_Sort "github.com/custergo/study_algo/3.Play-with-Algorithms/03-Sorting-Advance/03-Merge-Sort-Advance"
	"testing"
)

func TestQuickSort(t *testing.T) {
	n := 1000000
	arr := GenerateRandomArray(n, 0, n)
	//arr := GenerateNearlyOrderArray(n, 100)
	arr2 := make([]int, n)
	copy(arr2, arr)
	TestSort("MergeSort", Merge_Sort.MergeSort, arr)
	TestSort("QuickSort", QuickSort, arr2)
}
