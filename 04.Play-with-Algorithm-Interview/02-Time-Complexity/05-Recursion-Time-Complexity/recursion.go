package main

import "fmt"

// 递归中进行一次递归调用
// 二分查找算法的递归实现 Time Complexity O(logN)
func binarySearch(arr []int, l, r, target int) int {
	if l > r {
		return -1
	}
	mid := l + (r-l)/2
	if arr[mid] == target {
		return mid
	} else if arr[mid] > target {
		return binarySearch(arr, l, mid-1, target)
	} else {
		return binarySearch(arr, mid+1, r, target)
	}
}

// 计算从1加到n，递归实现 Time Complexity O(N)
func sum(n int) int {
	if n < 0 {
		panic("n不能小于0")
	}
	if n == 0 {
		return 0
	}
	return n + sum(n-1)
}

// 计算x的幂次方,递归实现 Time Complexity O(logN)
func pow(x float64, n int) float64 {
	if n < 0 {
		panic("n不能小于0")
	}

	if n == 0 {
		return 1.0
	}
	t := pow(x, n/2)
	if n%2 != 0 { // 为奇数
		return x * t * t
	}
	return t * t // 为偶数
}

// 递归中进行多次递归调用
// 		 计算调用次数
//            3
//       /        \
//      2          2
//    /  \        / \
//   1    1      1    1
//  /\    /\    /\    /\
// 0  0  0  0  0  0  0  0
func f(n int) int {
	if n < 0 {
		panic("n不得小于0")
	}
	if n == 0 {
		return 1
	}
	return f(n-1) + f(n-1)
}

//func mergeSort(arr []int, l, r int) {
//	if l >= r {
//		return
//	}
//	mid := l + (r-l)/2
//	mergeSort(arr, l, mid)
//	mergeSort(arr, mid+1, r)
//	merge(arr, l, mid, r)
//}

func main() {
	fmt.Println(sum(100))
	fmt.Println(pow(2, 10))
	fmt.Println(f(3))
}
