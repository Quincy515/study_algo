package Binary_Search

// 二分查找法,在有序数组arr中,查找target
// 如果找到target,返回相应的索引index
// 如果没有找到target就返回-1
func binarySearch(arr []int, n int, target int) int {
	// 在arr[l...r]之中查找target
	l, r := 0, n-1
	for l <= r {
		// 防止极端情况下的整型溢出,使用下面的逻辑求mid
		mid := l + (r-l)/2
		if arr[mid] == target {
			return mid
		}
		// 在arr[l...mid-1]之中查找target
		if arr[mid] > target {
			r = mid - 1
		} else { // target > arr[mid]
			// 在arr[mid+1...r]之中查找target
			l = mid + 1
		}
	}
	return -1
}

// 用递归方式写二分查找法
func binarySearchRec(arr []int, l, r, target int) int {
	if l > r {
		return -1
	}

	mid := l + (r-l)/2
	if arr[mid] == target {
		return mid
	} else if arr[mid] > target {
		return binarySearchRec(arr, l, mid-1, target)
	} else {
		return binarySearchRec(arr, mid+1, r, target)
	}
}

// 递归recursive
func binarySearchRecursive(arr []int, n, target int) int {
	return binarySearchRec(arr, 0, n-1, target)
}
