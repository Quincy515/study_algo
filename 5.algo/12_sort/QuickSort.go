package main

func QuickSort(arr []int) {
	separator(arr, 0, len(arr)-1)
}

func separator(arr []int, start, end int) {
	if start >= end {
		return
	}
	i := partition(arr, start, end)
	separator(arr, start, i-1)
	separator(arr, i+1, end)
}

func partition(arr []int, start, end int) int {
	// 选取最后一位当对比数字
	pivot := arr[end]
	var i = start
	for j := start; j < end; j++ {
		if arr[j] < pivot {
			if !(i == j) {
				//交换位置
				arr[i], arr[j] = arr[j], arr[i]
			}
			i++
		}
	}
	arr[i], arr[end] = arr[end], arr[i]
	return i
}
