package leetcode

// Time: O(n), Space: O(1)
func rob(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	n := len(nums)
	if n == 1 {
		return nums[0] // 返回抢这唯一一所房子的钱
	}
	return max(robb(nums, 0, n-2), robb(nums, 1, n-1))
}

// 辅助函数抢劫连排房子 输入是开始和结束下标 输出是最多抢到的钱
func robb(nums []int, start, end int) int {
	prev2, prev1 := 0, 0            // 初始化抢劫上上个房子的最大获利，和抢劫上个房子的最大获利
	for i := start; i <= end; i++ { // 从开始下标遍历到结束下标
		cur := max(prev1, prev2+nums[i])
		prev2 = prev1
		prev1 = cur
	}
	return prev1
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
