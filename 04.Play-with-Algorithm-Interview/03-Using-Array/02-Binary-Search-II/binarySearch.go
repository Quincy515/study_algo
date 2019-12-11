package main

func binarySearch(arr []interface{}, n int, target interface{}) int {
	l, r := 0, n // 边界设置的实际意义:在[l...r)的范围里寻找target
	for l < r {  // 当l==r时,区间[l...r)是无效的,如[42,42)
		mid := l + (r-l)/2
		if arr[mid] == target {
			return mid
		}
		if compare(target, arr[mid]) > 0 {
			l = mid + 1 // target在[mid+1...r)中
		} else { // target < arr[mid]
			r = mid // target在[l...mid)中
		}
	}
	return -1
}
