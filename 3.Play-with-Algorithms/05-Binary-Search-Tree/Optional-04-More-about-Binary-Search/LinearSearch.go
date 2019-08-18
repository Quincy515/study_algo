package Lower_Bound_and_Upper_Bound_in_Binary_Search

// 线性查找法,实现lower_bound
// 即在一个有序数组arr中,寻找大于等于target的元素的第一个索引
// 如果存在,则返回相应的索引index
// 否则,返回arr的元素个数n
func firstGreaterOrEqualLinearSearch(arr []int, target int) int {
	if arr == nil {
		panic("arr can not be null.")
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] >= target {
			return i
		}
	}
	return len(arr)
}

// 线性查找法,实现upper_bound
// 即在一个有序数组arr中,寻找大于target的元素的第一个索引
// 如果存在,则返回相应的索引index
// 否则,返回arr的元素个数n
func firstGreaterThanLinearSearch(arr []int, target int) int {
	if arr == nil {
		panic("arr can not be null.")
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] > target {
			return i
		}
	}
	return len(arr)
}

// 线性查找法,在一个有序数组arr中,寻找小于等于target元素的最大索引
// 如果存在,则返回相应的索引Index
// 否则,返回-1
func lastLessOrEqualLinearSearch(arr []int, target int) int {
	if arr == nil {
		panic("arr can not be null.")
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] > target {
			return i - 1
		}
	}
	return len(arr) - 1
}

// 线性查找法,在一个有序数组arr中,寻找小于target的元素的最大索引
// 如果存在,则返回相应的索引index
// 否则,返回-1
func lastLessThanLinearSearch(arr []int, target int) int {
	if arr == nil {
		panic("arr can not be null.")
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] >= target {
			return i - 1
		}
	}

	return len(arr) - 1
}
