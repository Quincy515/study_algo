package LeetCode

import "fmt"

// 自己思路
func minSubArrayLen1(s int, nums []int) int {
	count := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		count[i+1] = count[i] + nums[i]
	}
	fmt.Println(count)

	ret := len(nums) + 1
	for i := 0; i < len(nums); i++ {
		for k := len(nums) - i; k > 0; k-- {
			if count[i+k]-count[i] >= s {
				ret = min(ret, k)
			}
		}
	}
	if ret == len(nums)+1 {
		return 0
	}
	return ret
}

// 暴力解法 时间复杂度: O(n^3) 空间复杂度: O(1)-LeetCode超时
func minSubArrayLen2(s int, nums []int) int {
	res := len(nums) + 1
	for l := 0; l < len(nums); l++ {
		for r := l; r < len(nums); r++ {
			sum := 0
			for i := l; i <= r; i++ {
				sum += nums[i]
			}
			if sum >= s {
				res = min(res, r-l+1)
			}
		}
	}
	if res == len(nums)+1 {
		return 0
	}
	return res
}

// 优化暴力解 时间复杂度: O(n^2) 空间复杂度: O(1)
func minSubArrayLen3(s int, nums []int) int {
	sums := make([]int, len(nums)+1) // sums[i]存放nums[0...i-1]的和
	for i := 1; i <= len(nums); i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}

	res := len(nums) + 1
	for l := 0; l < len(nums); l++ {
		for r := l; r < len(nums); r++ {
			// 使用sums[r+1] - sums[l]快速获取nums[l...r]的和
			if sums[r+1]-sums[l] >= s {
				res = min(res, r-l+1)
			}
		}
	}

	if res == len(nums)+1 {
		return 0
	}
	return res
}

// 滑动窗口 时间复杂度: O(n) 空间复杂度: O(1)
func minSubArrayLen4(s int, nums []int) int {
	l, r := 0, -1        //nums[l...r]为我们的滑动窗口,如果[0,0]就包含了第一个元素,初始不包含任何元素
	sum := 0             // 合初始设置为0
	res := len(nums) + 1 // 最短数组长度设置为整个数组长度+1,设置为最大值

	for l < len(nums) { // 窗口的左边界在数组范围内,则循环继续
		if r+1 < len(nums) && sum < s {
			r++            // 窗口右边界扩大
			sum += nums[r] // 计算sum值
		} else { // r已经到头,或者sum>=s
			sum -= nums[l]
			l++ // 窗口左边界缩小
		}
		if sum >= s {
			res = min(res, r-l+1)
		}
	}
	if res == len(nums)+1 {
		return 0
	}
	return res
}

// 另一个滑动窗口的实现,仅供参考 时间复杂度: O(n) 空间复杂度: O(1)
func minSubArrayLen5(s int, nums []int) int {
	l, r := 0, -1 // [l...r] 为我们的窗口
	sum := 0
	res := len(nums) + 1

	for r+1 < len(nums) { // 窗口的右边界无法继续扩展了,则循环停止
		for r+1 < len(nums) && sum < s {
			r++
			sum += nums[r]
		}
		if sum >= s {
			res = min(res, r-l+1)
		}

		for l < len(nums) && sum >= s {
			sum -= nums[l]
			l++
			if sum >= s {
				res = min(res, r-l+1)
			}
		}
	}
	if res == len(nums)+1 {
		return 0
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 二分搜索,扩展优化暴力的方法.对于每一个l,可以使用二分搜索法搜索r
// 时间复杂度: O(nlogn) 空间复杂度: O(n)
func minSubArrayLen6(s int, nums []int) int {
	sums := make([]int, len(nums)+1) // sums[i]存放nums[0...i-1]的和
	for i := 1; i <= len(nums); i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}
	res := len(nums) + 1
	for l := 0; l < len(nums); l++ {
		// 实现一个基于二分搜索的lowerBound
		r := lowerBound(sums, sums[l]+s)
		if r != len(sums) {
			res = min(res, r-l)
		}
	}
	if res == len(nums)+1 {
		return 0
	}
	return res
}

// 在有序数组nums中寻找大于等于target的最小值
// 如果没有(nums数组中所有值都小于target),则返回nums.length
func lowerBound(nums []int, target int) int {
	if nums == nil /*|| !isSorted(nums)*/ {
		panic("Illegal argument nums in lowerBound.")
	}

	l, r := 0, len(nums) // 在nums[l...r)的范围内寻找解
	for l != r {
		mid := l + (r-l)/2
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func isSorted(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			return false
		}
	}
	return true
}
