package leetcode

// The: O(n), Space: O(n)
func rob(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0 // 数组为空或长度为0，直接返回0
	}
	if len(nums) == 1 {
		return nums[0] // 只有一个可以抢
	}
	// 否则说明数组中至少有两个元素
	d := make([]int, len(nums))                 // 定义状态数组d
	d[0], d[1] = nums[0], max(nums[0], nums[1]) // 初始条件
	// 从第2号房子开始遍历
	for i := 2; i < len(nums); i++ {
		d[i] = max(d[i-1], d[i-2]+nums[i])
	}
	return d[len(nums)-1] // 计算结束后返回最后的状态值即可
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// The: O(n), Space: O(1)
func robO1(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0 // 数组为空或长度为0，直接返回0
	}
	if len(nums) == 1 {
		return nums[0] // 只有一个可以抢
	}
	// 否则说明数组中至少有两个元素
	prev2, prev1 := nums[0], max(nums[0], nums[1]) // 表示抢到上上个房子的最大获利和抢到上个房子的最大获利
	// 从第2号房子开始遍历
	for i := 2; i < len(nums); i++ {
		cur := max(prev1, prev2+nums[i]) // 抢到当前房子的最大获利
		prev2 = prev1                    // 滚动更新抢到上上个房子的最大获利
		prev1 = cur                      // 滚动更新当前上个房子的最大获利
	}
	return prev1 // 计算结束后返回最后的状态值即可
}

// The: O(n), Space: O(1)
func robO11(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0 // 数组为空或长度为0，直接返回0
	}
	// 否则说明数组中至少有两个元素
	prev2, prev1 := 0, 0 // 表示抢到上上个房子的最大获利和抢到上个房子的最大获利
	// 从第2号房子开始遍历
	for _, num := range nums {
		cur := max(prev1, prev2+num) // 抢到当前房子的最大获利
		prev2 = prev1                // 滚动更新抢到上上个房子的最大获利
		prev1 = cur                  // 滚动更新当前上个房子的最大获利
	}
	return prev1 // 计算结束后返回最后的状态值即可
}
