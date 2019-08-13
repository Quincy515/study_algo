package Insertion_Sort

func InsertionSort(arr []int, n int) {
	for i := 1; i < n; i++ {
		// 寻找元素arr[i]合适的插入位置
		/* 写法一
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				swap(arr, arr[j], arr[j-1])
			} else {
				break
			}
		}
		*/
		// 写法二
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			swap(arr, j-1, j)
		}
	}
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
