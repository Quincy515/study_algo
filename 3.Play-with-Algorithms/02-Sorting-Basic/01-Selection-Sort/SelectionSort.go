package Selection_Sort

func SelectionSort(arr []int, n int) {
	for i := 0; i < n-1; i++ {
		// 寻找[i,n)区间里的最小值
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		swap(arr, i, minIndex)
	}
}

// slice 是引用类型
func swap(arr []int, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}
