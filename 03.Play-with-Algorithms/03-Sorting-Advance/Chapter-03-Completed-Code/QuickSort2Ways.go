package Sort

import (
	"math/rand"
)

// 双路快速排序对 arr[l...r] 部分进行 partition 操作
// 返回 p,使得 arr[l...p-1] < arr[p]; arr[p+1...r] > arr[p]
func partition2ways(arr []int, l, r int) int {
	// 随机在arr[l...r]的范围内，选择一个数值作为标的点pivot
	pivot := rand.Int()%(r-l+1) + l
	arr[l], arr[pivot] = arr[pivot], arr[l]
	v := arr[l]

	// arr[l+1...i] <= v; arr[j...r] >= v
	i, j := l+1, r
	for {
		// 思考,这里的边界,arr[i]<v,不能是arr[i]<=v
		for i <= r && arr[i] < v {
			i++
		}
		// 思考,这里的边界,arr[j]>v,不能是arr[j]>=v
		for j >= l+1 && arr[j] > v {
			j--
		}
		// 参考:http://coding.imooc.com/learn/questiondetail/4920.html
		// <=和>=会将重复元素归为一路，这样就不平衡
		if i > j {
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	arr[l], arr[j] = arr[j], arr[l]
	return j
}

// 对 arr[l...r]部分进行快速排序
func quickSort2ways(arr []int, l, r int) {
	//if l >= r {
	//	return
	//}
	// 对于小规模数组，使用插入排序
	if r-l <= 15 {
		insertionSort(arr, l, r)
		return
	}

	// 调用双路快速排序的partition
	p := partition2ways(arr, l, r)
	quickSort2ways(arr, l, p-1)
	quickSort2ways(arr, p+1, r)
}

func QuickSort2Ways(arr []int, n int) {
	quickSort2ways(arr, 0, n-1)
}
