package leetcode

// 使用辅助数组 Time: O(n), Space: O(n)
func minCostClimbingStairs(cost []int) int {
	if cost == nil || len(cost) == 0 {
		// 代价为空或长度为0
		return 0 // 直接返回0
	}
	if len(cost) == 1 {
		// 只有一阶楼梯，最小代价就是这阶楼梯对应的代价
		return cost[0]
	}
	n := len(cost)                // 先把代价数组的长度赋值给n
	d := make([]int, n)           // 定义长度为n的状态数组
	d[0], d[1] = cost[0], cost[1] // 初始化状态数组d(0)和d(1)
	for i := 2; i < n; i++ {      // 从第2阶楼梯向上爬
		// 爬到第i阶楼梯的最小代价为
		// 爬到前两阶楼梯最小代价的较小值
		// 再加上当前楼梯的代价
		d[i] = min(d[i-1], d[i-2]) + cost[i]
	}
	return min(d[n-1], d[n-2]) // 最后返回两个元素中较小值即可
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 使用两个变量滚动更新 Time: O(n), Space: O(n)
func minCostClimbingStairsO1(cost []int) int {
	if cost == nil || len(cost) == 0 {
		// 代价为空或长度为0
		return 0 // 直接返回0
	}
	if len(cost) == 1 {
		// 只有一阶楼梯，最小代价就是这阶楼梯对应的代价
		return cost[0]
	}
	first, second := cost[0], cost[1] // 初始化最初的两个最小代价
	for i := 2; i < len(cost); i++ {  // 从第2阶楼梯向上爬
		// 爬到第i阶楼梯的最小代价为
		// 爬到前两阶楼梯最小代价的较小值
		// 再加上当前楼梯的代价
		cur := min(first, second) + cost[i]
		// 滚动更新,到达前两阶楼梯的最小代价
		first = second
		second = cur
	}
	return min(first, second) // 循环结束后返回这两个代价中的较小值即可
}
