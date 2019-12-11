package Selection_Sort

import (
	"testing"
)

func TestSelectionSort(t *testing.T) {
	// 测试排序算法辅助函数
	N := 20000
	arr := GenerateRandomArray(N, 0, N)
	TestSort("SelectionSort", SelectionSort, arr)
	arr = nil
}
