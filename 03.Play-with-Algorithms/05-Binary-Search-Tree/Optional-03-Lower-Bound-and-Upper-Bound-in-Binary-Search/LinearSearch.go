package Lower_Bound_and_Upper_Bound_in_Binary_Search

// 线性查找法,实现lower_bound
// 即在一个有序数组arr中,寻找大于等于target的元素的第一个索引
// 如果存在,则返回相应的索引index
// 否则,返回arr的元素个数n
func lower_bound_linear_search(arr []int, target int) int {
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
func upper_bound_linear_search(arr []int, target int) int {
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
