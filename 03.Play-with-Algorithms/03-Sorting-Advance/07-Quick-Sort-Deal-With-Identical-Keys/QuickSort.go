package Quick_Sort

import (
	"math/rand"
)

// 双路快速排序
// 对 arr[l...r] 部分进行 partition 操作
// 返回 p,使得 arr[l...p-1] < arr[p]; arr[p+1...r] > arr[p]
func partition(arr []int, l, r int) int {
	// 优化2: 随机在arr[l...r]的范围内，选择一个数值作为标的点pivot
	pivot := rand.Int()%(r-l+1) + l
	arr[l], arr[pivot] = arr[pivot], arr[l]
	v := arr[l]

	// arr[l+1...i] <= v; arr[j...r] >= v
	i, j := l+1, r
	for {
		for i <= r && arr[i] < v {
			i++
		}
		for j >= l+1 && arr[j] > v {
			j--
		}
		if i > j {
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	arr[l], arr[j] = arr[j], arr[l]
	return j
}

// 对 arr[l...r]部分进行快速排序
func quickSort(arr []int, l, r int) {
	//if l >= r {
	//	return
	//}
	// 优化1: 对于小规模数组，使用插入排序
	if r-l <= 15 {
		insertionSort(arr, l, r)
		return
	}

	p := partition(arr, l, r)
	quickSort(arr, l, p-1)
	quickSort(arr, p+1, r)
}

func QuickSort(arr []int, n int) {
	quickSort(arr, 0, n-1)
}

func insertionSort(arr []int, l, r int) {
	for i := l; i <= r; i++ {
		e := arr[i] // 寻找元素arr[i]合适的插入位置
		var j int   // j保存元素temp应该插入的位置
		for j = i; j > l && arr[j-1] > e; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = e
	}
}
