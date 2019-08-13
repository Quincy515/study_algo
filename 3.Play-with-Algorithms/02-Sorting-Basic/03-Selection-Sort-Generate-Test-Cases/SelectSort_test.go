package Selection_Sort

import (
	"fmt"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	// 测试排序算法辅助函数
	N := 20000
	arr := GenerateRandomArray(N, 0, N)
	SelectionSort(arr, len(arr))
	fmt.Println(arr)
	arr = nil
}
