package Binary_Search

import (
	"fmt"
	"testing"
	"time"
)

// 比较非递归和递归写法的二分查找的效率
// 非递归算法在性能上有微弱的优势
func TestBinarySearch(t *testing.T) {
	n := 1000000
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = i
	}

	// 测试非递归二分查找法
	startTime := time.Now()
	// 对于待查找数组[0...N)
	// 对[0...N)区间的数值使用二分查找,最终结果应该就是数字本省
	// 堆[N...2*N)区间的数值使用二分查找,因为这些数字不在arr中,结果为-1
	for i := 0; i < 2*n; i++ {
		v := binarySearch(a, n, i)
		if i < n {
			if v != i {
				panic("查找失败!")
			}
		} else {
			if v != -1 {
				panic("查找失败2!")
			}
		}
	}
	fmt.Println("Binary Search(Without Recursion): ", time.Now().Sub(startTime))

	// 测试递归的二分查找法
	startTime = time.Now()
	for i := 0; i < 2*n; i++ {
		v := binarySearchRecursive(a, n, i)
		if i < n {
			if v != i {
				panic("查找失败!!")
			}
		} else {
			if v != -1 {
				panic("查找失败2!" +
					"1")
			}
		}
	}
	fmt.Println("Binary Search(Recursion): ", time.Now().Sub(startTime))
}
