package Merge_Sort_Bottom_Up

import "math"

// 将arr[l...mid]和arr[mid+1...r]两部分归并
func merge(arr []int, l, mid, r int) {
	aux := make([]int, r-l+1)
	for i := l; i <= r; i++ {
		aux[i-l] = arr[i]
	}

	// 初始化，i指向左半部分的起始索引位置i；j指向右半部分起始索引位置mid+1
	i, j := l, mid+1
	for k := l; k <= r; k++ {
		if i > mid { // 如果左半部分元素已经全部处理完毕
			arr[k] = aux[j-l]
			j++
		} else if j > r { // 如果右半部分元素已经全部处理完毕
			arr[k] = aux[i-l]
			i++
		} else if aux[i-l] < aux[j-l] { // 左半部分所指元素 < 右半部分所指元素
			arr[k] = aux[i-l]
			i++
		} else { // 左半部分所指元素 >= 右半部分所指元素
			arr[k] = aux[j-l]
			j++
		}
	}
}

func MergeSortBU(arr []int, n int) {
	for size := 1; size <= n; size += size {
		for i := 0; i+size < n; i += size * 2 {
			// 对arr[i...i+size-1]和arr[i+size...i+2*size-1]进行归并
			merge(arr, i, i+size-1, int(math.Min(float64(i+size*2-1), float64(n-1))))
		}
	}
}
