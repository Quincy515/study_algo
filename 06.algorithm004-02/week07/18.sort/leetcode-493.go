package leetcode

// 暴力求解超时 Time: O(n^2),Space: O(1)
func reversePairsForce(nums []int) int {
	if nums == nil || len(nums) < 2 {
		return 0 // 数组为空，或长度小于2，返回0
	}
	cnt := 0 // 定义计数器
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > 2*nums[j] {
				cnt++
			}
		}
	}
	return cnt
}

// 辅助函数 合并两个有序子数组
func merge(arr []int, low, mid, high int, tmp []int) {
	i, j, k := low, mid+1, 0
	for i <= mid && j <= high {
		if arr[i] <= arr[j] {
			tmp[k] = arr[i]
			k, i = k+1, i+1
		} else {
			tmp[k] = arr[j]
			k, j = k+1, j+1
		}
	}
	for i <= mid {
		tmp[k] = arr[i]
		k, i = k+1, i+1
	}
	for j <= high {
		tmp[k] = arr[j]
		k, j = k+1, j+1
	}
	copy(arr[low:high+1], tmp)
}

// 辅助函数 统计重要逆序对
func count(arr []int, low, mid, high int) int {
	i, j := low, mid+1
	cnt := 0
	for i <= mid && j <= high {
		if arr[i] <= 2*arr[j] {
			i++
		} else {
			cnt += (mid - i + 1)
			j++
		}
	}
	return cnt
}

// 辅助函数 归并排序和统计重要逆序对数量
func sortAndCount(arr []int, low, high int, tmp []int) int {
	cnt := 0
	if low < high {
		mid := low + (high-low)/2
		cnt += sortAndCount(arr, low, mid, tmp)
		cnt += sortAndCount(arr, mid+1, high, tmp)
		cnt += count(arr, low, mid, high)
		merge(arr, low, mid, high, tmp)
	}
	return cnt
}

// 归并排序 Time: O(n*log(n)), Space: O(n)
func reversePairs(nums []int) int {
	if nums == nil || len(nums) < 2 {
		return 0
	}
	tmp := make([]int, len(nums))
	return sortAndCount(nums, 0, len(nums)-1, tmp)
}
