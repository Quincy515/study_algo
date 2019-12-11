package Lower_Bound_and_Upper_Bound_in_Binary_Search

// 二分查找法,实现lower_bound
// 即在一个有序数组arr中，寻找大于等于target的元素的第一个索引
// 如果存在,则返回相应的索引index
// 否则,返回arr的元素个数n
func lower_bound(arr []int, target int) int {
	if arr == nil {
		panic("arr can not be null.")
	}

	l, r := 0, len(arr)
	for l != r {
		mid := l + (r-l)/2
		if arr[mid] < target {
			l = mid + 1
		} else { // arr[mid] >= target
			r = mid
		}
	}
	return l
}

// 二分查找法,实现upper_bound
// 即在一个有序数组arr中,寻找大于target的元素的第一个索引
// 如果存在,则返回相应的索引index
// 否则,返回arr的元素个数n
func upper_bound(arr []int, target int) int {
	if arr == nil {
		panic("arr can not be null.")
	}

	l, r := 0, len(arr)
	for l != r {
		mid := l + (r-l)/2
		if arr[mid] <= target {
			l = mid + 1
		} else { // arr[min] > target
			r = mid
		}
	}
	return l
}
