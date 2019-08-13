package Insertion_Sort

func InsertionSort(arr []int, n int) {
	for i := 1; i < n; i++ {
		e := arr[i] // 寻找元素arr[i]合适的插入位置
		var j int   // j保存元素temp应该插入的位置
		for j = i; j > 0 && arr[j-1] > e; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = e
	}
}
