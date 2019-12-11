package Sort

func ShellSort(arr []int, length int) {
	h := 1
	for h < length/3 {
		h = 3*h + 1
	}
	// 计算 increment sequence: 1, 4, 13, 40, 121, 364, 1093...

	for h >= 1 {
		// h-sort the array
		for i := h; i < length; i++ {
			// 对 arr[i], arr[i-h], arr[i-2*h], arr[i-3*h]... 使用插入排序
			e := arr[i]
			j := i
			for ; j >= h && e < arr[j-h]; j -= h {
				arr[j] = arr[j-h]
			}
			arr[j] = e
		}
		h /= 3
	}
}
