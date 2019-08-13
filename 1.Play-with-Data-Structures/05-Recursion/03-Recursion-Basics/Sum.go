package _3_Recursion_Basics

func Sum(arr []int) int {
	return sum(arr, 0)
}

// 计算arr[l...n)这个区间内所有数字的和
func sum(arr []int, l int) int {
	// 求解最基本问题
	if l == len(arr) {
		return 0
	}
	// 把原问题转换成更小的问题
	return arr[l] + sum(arr, l+1)
}
