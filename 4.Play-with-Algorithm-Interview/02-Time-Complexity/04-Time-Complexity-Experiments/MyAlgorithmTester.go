package main

// O(logN) 二分查找法
func binarySearch(arr []int, n, target int) int {
	l, r := 0, n-1
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

// O(N) 在数组中寻找最大值
func findMax(arr []int, n int) int {
	assert(n > 0)
	res := arr[0]
	for i := 1; i < n; i++ {
		if arr[i] > res {
			res = arr[i]
		}
	}
	return res
}

// O(NlogN) 归并排序
func _merge(arr []int, l, mid, r int, aux []int) {
	i, j := l, mid+1
	for k := l; k <= r; k++ {
		if i > mid {
			arr[k] = aux[k]
			j++
		} else if j > r {
			arr[k] = aux[i]
			i++
		} else if aux[i] < aux[j] {
			arr[k] = aux[i]
			i++
		} else {
			arr[k] = aux[j]
			j++
		}
	}
}

// 自底向上的归并排序,所以不需要使用递归
func mergeSort(arr []int, n int) {
	aux := make([]int, n)
	for i := 0; i < n; i++ {
		aux[i] = arr[i]
	}

	for sz := 1; sz < n; sz += sz {
		for i := 0; i < n; i += sz + sz {
			_merge(arr, i, i+sz-1, min(i+sz+sz-1, n-1), aux)
		}
	}
}

// O(N^2) 选择排序
func selectionSort(arr []int, n int) {
	for i := 0; i < n; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		swap(arr, i, minIndex)
	}
}
