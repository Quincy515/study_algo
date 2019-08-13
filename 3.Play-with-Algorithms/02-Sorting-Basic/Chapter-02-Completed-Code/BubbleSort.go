package Sort

func BubbleSort(arr []int, length int) {
	for i := 0; i < length-1; i++ {
		for j := 1; j < length-i; j++ {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[i]
			}
		}
	}
}
