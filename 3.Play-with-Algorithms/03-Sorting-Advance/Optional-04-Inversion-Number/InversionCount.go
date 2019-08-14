package Inversion_Number

// 有一组数，对于其中任意两个数字，若前面一个大于后面一个数字，则这两个数字组成一个逆序对
// [1,2,3,4,5,6,7,0] 逆序对个数: 7
// 对于一个大小为N的数组, 其最大的逆序数对个数为N*(N-1)/2,非常容易产生整型溢出,使用long long返回

// merge函数求出在arr[l...mid]和arr[mid+1...r]有序的基础上,arr[l...r]的逆序数对个数
func merge(arr []int, l, mid, r int) int64 {
	aux := make([]int, r-l+1)
	for i := l; i <= r; i++ {
		aux[i-l] = arr[i]
	}

	// 初始化逆序数对个数res =0
	res := int64(0)
	// 初始化,i指向左半部分的起始索引位置l,j指向右半部分起始索引位置mid+1
	j, k := l, mid+1
	for i := l; i <= r; i++ {
		if j > mid { // 如果左半部分元素已经全部处理完毕
			arr[i] = aux[k-l]
			k++
		} else if k > r { // 如果右半部分元素已经全部处理完毕
			arr[i] = aux[j-l]
			j++
		} else if aux[j-l] <= aux[k-l] { // 左半部分所指元素 <= 右半部分所指元素
			arr[i] = aux[j-l]
			j++
		} else { // 右半部分所指元素 < 左半部分所指元素
			arr[i] = aux[k-l]
			k++
			// 此时, 因为右半部分k所指的元素小
			// 这个元素和左半部分的所有未处理的元素都构成了逆序数对
			// 左半部分此时未处理的元素个数为mid-j+1
			res += int64(mid - j + 1)
		}
	}
	return res
}

// 求 arr[l...r]范围的逆序数对个数
// 思考: 归并排序的优化可否用于求逆序数对的算法?
func inversionCount(arr []int, l, r int) int64 {
	if l >= r {
		return 0
	}
	mid := l + (r-l)/2
	// 求出arr[l...mid]范围的逆序数
	res1 := int64(inversionCount(arr, l, mid))
	// 求出arr[mid+1...r]范围的逆序数
	res2 := int64(inversionCount(arr, mid+1, r))
	return res1 + res2 + merge(arr, l, mid, r)
}

// 递归求arr的逆序对个数

func InversionCount(arr []int, n int) int64 {
	return inversionCount(arr, 0, n-1)
}
