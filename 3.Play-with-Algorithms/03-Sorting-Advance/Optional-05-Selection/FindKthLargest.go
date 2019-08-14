package Optional_05_Selection

import "math/rand"

// partition过程,和快排的partition一样
// 思考: 双路快排和三路快排的思想能不能用在selection算法中?
func partition(arr []int, l, r int) int {
	p := rand.Int()%(r-l+1) + l
	swap(arr, l, p)

	j := l // [l+1...j]<p; [lt...i]>p
	for i := l + 1; i <= r; i++ {
		if arr[i] < arr[l] {
			j++
			swap(arr, i, j)
		}
	}
	swap(arr, l, j)
	return j
}

// 求出arr[l...r]范围里第k小的数
func selection(arr []int, l, r, k int) int {
	if l == r {
		return arr[l]
	}
	// partition之后,arr[p]的正确位置就在索引p上
	p := partition(arr, l, r)

	if k == p { // 如果k==p,直接返回arr[p]
		return arr[p]
	} else if k < p { // 如果k<p,只需要在arr[l...p-1]中找到第k小元素即可
		return selection(arr, l, p-1, k)
	} else { // 如果k>p,则需要在arr[p+1...r]中找出第k-p-1小元素
		// 注意: 由于传入selection的依然是arr,而不是arr[p+1...r]
		// 所以传入的最后一个参数依然是k,而不是k-p-1
		return selection(arr, p+1, r, k)
	}
}

// 寻找arr数组中第k小的元素
// 注意: 在算法中, k是从0开始索引的, 即最小的元素时第0小元素, 以此类推
// 如果希望算法中k的语意是从1开始的, 只需要在整个逻辑开始进行k--即可, 可以参考 selection2
func Selection(arr []int, n, k int) int {
	if k < 0 || k > n {
		panic("Error")
	}
	return selection(arr, 0, n-1, k)
}

// 寻找arr数组中第k小的元素, k从1开始索引, 即最小元素是第1小元素, 以此类推
func Selection2(arr []int, n, k int) int {
	return Selection(arr, n, k-1)
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
