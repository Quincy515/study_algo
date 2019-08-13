package Bubble_Sort

func BubbleSort1(arr []int, length int) {
	var swapped bool
	for {
		swapped = false
		for i := 1; i < length; i++ {
			if arr[i] < arr[i-1] {
				arr[i], arr[i-1] = arr[i-1], arr[i]
				swapped = true
			}
		}
		// 循环中没有数据交换，说明已排序，停止循环
		if !swapped {
			break
		}
	}
}
