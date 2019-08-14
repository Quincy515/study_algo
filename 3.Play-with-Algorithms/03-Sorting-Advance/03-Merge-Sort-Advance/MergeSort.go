package Merge_Sort

// 将arr[l...mid]和arr[mid_1...r]两部分归并
func merge(arr []int, l, mid, r int) {
	aux := make([]int, r-l+1)
	for i := l; i <= r; i++ {
		aux[i-l] = arr[i]
	}

	// 初始化，i指向左半部分的起始索引位置l；j指向右半部分起始索引位置mid+1
	i, j := l, mid+1
	for k := l; k <= r; k++ {
		if i > mid { // 如果左部分元素已经全部处理完毕
			arr[k] = aux[j-l]
			j++
		} else if j > r { // 如果右半部分元素已经全部处理完毕
			arr[k] = aux[i-l]
			i++
		} else if aux[i-l] < aux[j-l] { // 左半部分所指元素 < 右半部分所指元素
			arr[k] = aux[i-l]
			i++
		} else {
			arr[k] = aux[j-l]
			j++
		}
	}
}

// 递归使用归并排序，对arr[l...r]的范围进行排序
func mergeSort(arr []int, l, r int) {
	//if l >= r {
	//	return
	//}
	// 优化2: 对于小规模数组，使用插入排序
	if r-l <= 15 {
		insertionSort(arr, l, r)
		return
	}
	mid := (l + r) / 2
	// 拆分的过程
	mergeSort(arr, l, mid)
	mergeSort(arr, mid+1, r)
	// 合并的过程
	// 优化1: 对于arr[mid]<=arr[mid+1]的情况，不进行merge
	// 对于近乎有序的数组非常有效，但是对于一般情况，有一定的性能损失
	if arr[mid] > arr[mid+1] {
		merge(arr, l, mid, r)
	}
}

func MergeSort(arr []int, n int) {
	mergeSort(arr, 0, n-1)
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
