package Quick_Sort

import (
	Merge_Sort "github.com/custergo/study_algo/3.Play-with-Algorithms/03-Sorting-Advance/03-Merge-Sort-Advance"
	"testing"
)

func TestQuickSort(t *testing.T) {
	n := 1000000
	arr := GenerateRandomArray(n, 0, n)
	nearlyarr := GenerateNearlyOrderArray(n, 100)
	repeatarr := GenerateRandomArray(n, 0, 10)
	arr2 := make([]int, n)
	nearlyarr2 := make([]int, n)
	repeatarr2 := make([]int, n)
	copy(arr2, arr)
	copy(nearlyarr2, nearlyarr)
	copy(repeatarr2, repeatarr)
	TestSort("MergeSort", Merge_Sort.MergeSort, arr)
	TestSort("QuickSort", QuickSort, arr2)

	TestSort("MergeSort", Merge_Sort.MergeSort, nearlyarr)
	TestSort("QuickSort", QuickSort, nearlyarr2)

	TestSort("MergeSort", Merge_Sort.MergeSort, repeatarr)
	TestSort("QuickSort", QuickSort, repeatarr2)
}
