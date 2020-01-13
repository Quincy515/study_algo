package leetcode

// TIme: O(n), Space: O(1)
func maxSubArray(nums []int) int {
	max := -1 << 32                  // 初始化连续子序列的最大和为int的最小值
	cur := 0                         // 初始化子序列当前的累加和为0
	for i := 0; i < len(nums); i++ { // 遍历数组
		if cur <= 0 { // 如果当前累加和小于等于0
			cur = nums[i] // 则用当前元素值nums[i]更新cur，表示新的子序列
		} else {
			cur += nums[i] // 否则继续累加当前元素的值
		}
		max = maximum(max, cur) // 更新当前子序列的最大值
	}
	return max // 循环结束后直接返回max即可
}

func maximum(a, b int) int {
	if a < b {
		return b
	}
	return a
}
