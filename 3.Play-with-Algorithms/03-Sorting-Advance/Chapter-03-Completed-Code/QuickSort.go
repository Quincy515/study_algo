package Sort

import (
	"math/rand"
)

// 对 arr[l...r] 部分进行 partition 操作
// 返回 p,使得 arr[l...p-1] < arr[p]; arr[p+1...r] > arr[p]
func partition(arr []int, l, r int) int {
	// 随机在arr[l...r]的范围内，选择一个数值作为标的点pivot
	pivot := rand.Int()%(r-l+1) + l
	arr[l], arr[pivot] = arr[pivot], arr[l]
	v := arr[l]

	// arr[l+1...j] < v; arr[j+1...i) > v
	j := l
	for i := l + 1; i <= r; i++ {
		if arr[i] < v {
			arr[j+1], arr[i] = arr[i], arr[j+1]
			j++
		}
	}
	arr[l], arr[j] = arr[j], arr[l]
	return j
}

// 对 arr[l...r]部分进行快速排序
func quickSort(arr []int, l, r int) {
	//if l >= r {
	//	return
	//}
	// 对于小规模数组，使用插入排序
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
