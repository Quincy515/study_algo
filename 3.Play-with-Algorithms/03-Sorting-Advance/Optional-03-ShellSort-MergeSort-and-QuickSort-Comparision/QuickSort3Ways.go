package Sort

import "math/rand"

// 三路快速排序处理arr[l...r] 将 arr[l...r] 分为 <v; ==v; >v 三部分
// 之后递归对 <v; >v 两部分继续进行三路快速排序
func quickSort3Ways(arr []int, l, r int) {
	// 对于小规模数组，使用插入排序
	if r-l <= 15 {
		insertionSort(arr, l, r)
		return
	}
	// partition,随机在arr[l...r]的范围中,选择一个数值作为标定点pivot
	swap(arr, l, rand.Int()%(r-l+1)+l)
	v := arr[l]

	lt := l     // arr[l+1...lt] < v
	gt := r + 1 // arr[gt...r]   > v
	i := l + 1  // arr[lt+1...i) == v
	for i < gt {
		if arr[i] < v {
			swap(arr, i, lt+1)
			lt++
			i++
		} else if arr[i] > v {
			swap(arr, i, gt-1)
			gt--
		} else { //arr[i] == v
			i++
		}
	}
	swap(arr, l, lt)
	quickSort3Ways(arr, l, lt-1)
	quickSort3Ways(arr, gt, r)
}

func QuickSort3Ways(arr []int, n int) {
	quickSort3Ways(arr, 0, n-1)
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
