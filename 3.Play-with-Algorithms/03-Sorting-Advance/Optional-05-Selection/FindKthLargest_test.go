package Optional_05_Selection

import (
	"fmt"
	"testing"
)

func TestSelection(t *testing.T) {
	// 生成一个大小为n,包含0...n-1这n个元素的随机数组arr
	n := 10000
	arr := GenerateOrderArray(n)
	ShuffleArray(arr, n)
	// 验证Selection算法,对arr数组求第i个元素,应该为i
	for i := 0; i < n; i++ {
		if Selection(arr, n, i) != i {
			panic("Selection算法错误!")
		}
	}
	fmt.Println("Test Selection Completed.")

	// 验证Selection2算法
	arr = GenerateOrderArray(n)
	ShuffleArray(arr, n)

	// 对arr数组求第i个元素,应该为i-1(在Selection2中,第k小元素的k是从1索引的)
	for i := 1; i <= n; i++ {
		if Selection2(arr, n, i) != i-1 {
			panic("Selection2算法错误!")
		}
	}
	fmt.Println("Test Selection2 Completed.")
}
