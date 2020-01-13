package leetcode

import "fmt"

// Time: O(n), Space: O(1)
func maxProduct(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0 // 如果数组为空，长度为0，则返回0
	}
	globalMax, curMax, curMin := nums[0], nums[0], nums[0] // 定义全局最大乘积，当前最大乘积，当前最小乘积
	for i := 1; i < len(nums); i++ {                       // 从数组第一个元素开始遍历数组
		a := curMax * nums[i]                  // 定义a是当前最大乘积乘以当前整数
		b := curMin * nums[i]                  // 定义b是当前最小乘积乘以当前整数
		c := nums[i]                           // 定义c是当前整数
		curMax = max(a, b, c)                  // 当前最大乘积定义为a,b,c中最大值
		curMin = min(a, b, c)                  // 当前最小乘积定义为a,b,c中最小值
		globalMax = maximum(curMax, globalMax) // 更新全局最大乘积
		fmt.Println("第 ", i, " 次,", "当前最大乘积: ", curMax, " 当前最小乘积: ", curMin, " 全局最大乘积: ", globalMax)
	}
	return globalMax // 循环结束后返回全局最大乘积即可
}

// 辅助函数 用于求三个整数的最大值和最小值
func max(a, b, c int) int {
	return maximum(maximum(a, b), c)
}

func min(a, b, c int) int {
	return minimum(minimum(a, b), c)
}

func maximum(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func minimum(a, b int) int {
	if a < b {
		return a
	}
	return b
}
