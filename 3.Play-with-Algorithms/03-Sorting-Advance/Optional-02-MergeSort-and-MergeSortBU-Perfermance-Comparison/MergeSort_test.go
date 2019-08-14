package Merge_Sort

import (
	"testing"
)

func TestMergeSort(t *testing.T) {
	n := 50000
	//arr := GenerateRandomArray(n, 0, n)
	arr := GenerateNearlyOrderArray(n, 10)
	arr2 := make([]int, n)
	arr3 := make([]int, n)
	arr4 := make([]int, n)
	copy(arr2, arr)
	copy(arr3, arr)
	copy(arr4, arr)
	TestSort("自顶向下的归并排序,无优化", MergeSort1, arr)
	TestSort("自顶向下的归并排序,含优化", MergeSort2, arr2)
	TestSort("自底向上的归并排序,无优化", MergeSortBU1, arr3)
	TestSort("自底向上的归并排序,含优化", MergeSortBU2, arr4)
}
