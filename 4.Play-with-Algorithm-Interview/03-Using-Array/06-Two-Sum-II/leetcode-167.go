package LeetCode

// 暴力枚举法 时间复杂度: O(n^2) 空间复杂度: O(1)
func twoSum1(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				return []int{i + 1, j + 1}
			}
		}
	}
	return []int{-1, -1}
}

// 二分搜索法 时间复杂度: O(nlogn) 空间复杂度: O(1)
func twoSum2(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		j := binarySearch(numbers, i+1, len(numbers)-1, target-numbers[i])
		if j != -1 {
			return []int{i + 1, j + 1}
		}
	}
	return []int{-1, -1}
}

func binarySearch(nums []int, l, r, target int) int {
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		}
		if target > nums[mid] {
			l = mid + 1
		} else { // target < nums[mid]
			r = mid - 1
		}
	}
	return -1
}

// 对撞指针 时间复杂度: O(n) 空间复杂度: O(1)
func twoSum3(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for i < j {
		sum := numbers[i] + numbers[j]
		if sum == target {
			return []int{i + 1, j + 1}
		} else if sum < target {
			i++
		} else { // sum < target
			j--
		}
	}
	return []int{-1, -1}
}
