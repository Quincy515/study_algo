package Sort

import (
	"fmt"
	"testing"
	"time"
)

// 比较 Shell Sort 和 Merge Sort 和 三种Quick Sort 的性能效率
// 使用更科学的比较方式,每次比较都运行多次测试用例，取平均值
// 可以看出, Shell Sort虽然鳗鱼高级的排序方式,但仍然是非常有竞争力的一种排序算法
// 其所花费的时间完全在可以容忍的范围内,远不像O(n^2)的排序算法,在数据量较大的时候无法忍受
// 同时,Shell Sort实现简单,只使用循环的方式解决排序问题,不需要实现递归,不占用系统空间,也不依赖随机数
// 所以,如果算法实现所使用的环境不利于实现复杂的排序算法,或者在项目工程的测试阶段,完全可以暂时使用ShellSort来进行排序任务:)
func TestQuickSort(t *testing.T) {
	// 测试T个测试用例，每个测试用例的数组大小为n
	T := 100
	n := 1000000
	var (
		time1 time.Duration
		time2 time.Duration
		time3 time.Duration
		time4 time.Duration
		time5 time.Duration
	)
	// 比较ShellSort和MergeSort和三种QuickSort的性能效率
	for i := 0; i < T; i++ {
		arr1 := GenerateRandomArray(n, 0, n)
		arr2 := make([]int, n)
		copy(arr2, arr1)
		arr3 := make([]int, n)
		copy(arr3, arr1)
		arr4 := make([]int, n)
		copy(arr4, arr1)
		arr5 := make([]int, n)
		copy(arr5, arr1)

		time1 += TestSortTime(ShellSort, arr1)
		time2 += TestSortTime(MergeSort, arr2)
		time3 += TestSortTime(QuickSort, arr3)
		time4 += TestSortTime(QuickSort2Ways, arr4)
		time5 += TestSortTime(QuickSort3Ways, arr5)
	}
	fmt.Println("ShellSort: ", time1)
	fmt.Println("MergeSort: ", time2)
	fmt.Println("QuickSort: ", time3)
	fmt.Println("QuickSort2Ways: ", time4)
	fmt.Println("QuickSort3Ways: ", time5)
}
