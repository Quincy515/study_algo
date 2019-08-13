package Bubble_Sort

func BubbleSort2(arr []int, length int) {
	var newn int
	for {
		newn = 0
		for i := 1; i < length; i++ {
			if arr[i] < arr[i-1] {
				arr[i], arr[i-1] = arr[i-1], arr[i]
				newn = i
			}
		}
		// 最后一次交换的点，意味着后面的值都正序，不需要再遍历
		length = newn
		// for 循环结束没有数值交换说明已排序
		if newn == 0 {
			break
		}
	}
}
