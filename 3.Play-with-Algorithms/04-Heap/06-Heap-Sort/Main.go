package Heap_Sort

func shiftDown(arr []int, n, k int) {
	for 2*k+1 < n { // 左孩子不越界
		// 默认为左孩子,在此轮循环中,arr[k]和arr[j]交换位置
		j := 2*k + 1
		// 右孩子不越界,右孩子大于左孩子
		if j+1 < n && arr[j+1] > arr[j] {
			j += 1 // 改变j为右孩子
		}
		if arr[k] >= arr[j] {
			break // 父亲节点大于子节点,满足二叉堆特性退出
		}
		arr[k], arr[j] = arr[j], arr[k] // k,j位置交换元素
		k = j                           // 继续查看该节点是否满足二叉堆特性
	}
}

func HeapSort(arr []int, n int) {
	for i := n - 1/2; i >= 0; i-- {
		shiftDown(arr, n, i)
	}

	for i := n - 1; i > 0; i-- {
		swap(arr, 0, i)
		shiftDown(arr, i, 0)
	}
}
