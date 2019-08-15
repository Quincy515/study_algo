package Heap_Sort

func HeapSort1(arr []int, n int) {
	maxHeap := NewMaxHeap(n)
	for i := 0; i < n; i++ {
		maxHeap.Insert(arr[i])
	}

	// 从小到大排序
	for i := n - 1; i >= 0; i-- {
		arr[i] = maxHeap.ExtractMax()
	}
}

func HeapSort2(arr []int, n int) {
	maxHeap := NewHeapify(arr, n)
	// 从小到大排序
	for i := n - 1; i >= 0; i-- {
		arr[i] = maxHeap.ExtractMax()
	}
}
