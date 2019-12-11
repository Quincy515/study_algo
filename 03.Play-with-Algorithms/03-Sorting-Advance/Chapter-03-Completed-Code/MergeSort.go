package Sort

import "math"

/************************
 * 自顶向下的归并排序,无优化 *
 ************************/

// 将arr[l...mid]和arr[mid+1...r]两部分归并
func merge1(arr []int, l, mid, r int) {
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

// 递归使用归并排序, 对arr[l...r]的范围进行排序
func mergeSort1(arr []int, l, r int) {
	if l >= r {
		return
	}
	mid := (l + r) / 2
	mergeSort1(arr, l, mid)
	mergeSort1(arr, mid+1, r)
	merge1(arr, l, mid, r)
}

// 自顶向下的归并排序, 无优化
func MergeSort1(arr []int, n int) {
	mergeSort1(arr, 0, n-1)
}

/************************
 * 自低向上的归并排序,无优化 *
 ************************/
func MergeSortBU1(arr []int, n int) {
	for size := 1; size <= n; size += size {
		for i := 0; i+size < n; i += size * 2 {
			// 对arr[i...i+size-1]和arr[i+size...i+2*size-1]进行归并
			merge1(arr, i, i+size-1, int(math.Min(float64(i+size*2-1), float64(n-1))))
		}
	}
}

/************************
 * 自顶向下的归并排序,含优化 *
 ************************/
func merge2(arr, aux []int, l, mid, r int) {
	// 由于aux的大小和arr一样,所以不需要处理aux索引的偏移量,进一步节省计算量
	i, j := l, mid+1
	for k := l; k <= r; k++ {
		if i > mid { // 如果左半部分元素已经全部处理完毕
			arr[k] = aux[j]
			j++
		} else if j > r { // 如果右半部分元素已经全部处理完毕
			arr[k] = aux[i]
			i++
		} else if aux[i] < aux[j] { // 左半部分所指元素 < 右半部分所指元素
			arr[k] = aux[i]
			i++
		} else { // 左半部分所指元素 >= 右半部分所指元素
			arr[k] = aux[j]
			j++
		}
	}
}

// 使用优化的归并排序算法,对arr[l...r]的范围进行排序
// 其中aux为完成merge过程所需要的辅助空间
func mergeSort2(arr, aux []int, l, r int) {
	// 对于小规模数组，使用插入排序
	if r-l >= 15 {
		insertionSort(arr, l, r)
		return
	}

	mid := (l + r) / 2
	mergeSort2(arr, aux, l, mid)
	mergeSort2(arr, aux, mid+1, r)

	// 对于arr[mid]<=arr[mid+1]的情况,不进行merge
	// 对于近乎有序的数组非常有效,但是对于一般情况,有一定的性能损失
	if arr[mid] > arr[mid+1] {
		merge2(arr, aux, l, mid, r)
	}
}

// 自顶向下的归并排序,含优化
func MergeSort2(arr []int, n int) {
	// 一次性申请aux空间,并将这个辅助空间以参数形式传递给完成归并排序的各个子函数
	aux := make([]int, n)
	mergeSort2(arr, aux, 0, n-1)
}

/************************
 * 自底向上的归并排序,含优化 *
 ************************/
func MergeSortBU2(arr []int, n int) {
	// 对于小规模数组，直接插入排序
	for i := 0; i < n; i += 16 {
		insertionSort(arr, i, int(math.Min(float64(i+15), float64(n-1))))
	}

	// 一次性申请aux空间,并将这个辅助空间以参数形式传递给完成归并排序的各个子函数
	aux := make([]int, n)
	for size := 16; size <= n; size += size {
		for i := 0; i < n-size; i += size + size {
			// 对于arr[mid]<=arr[mid+1]的情况,不进行merge
			// 对于近乎有序的数组非常有效,但是对于一般情况,有一定的性能损失
			if arr[i+size-1] > arr[i+size] {
				merge2(arr, aux, i, i+size-1, int(math.Min(float64(i+size+size-1), float64(n-1))))
			}
		}
	}
}

func insertionSort(arr []int, l, r int) {
	for i := l; i <= r; i++ {
		e := arr[i] // 寻找元素arr[i]合适的插入位置
		var j int   // j保存元素e应该插入的位置
		for j = i; j > l && arr[j-1] > e; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = e
	}
}
