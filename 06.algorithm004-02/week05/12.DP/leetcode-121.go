package leetcode

// Time: O(n^2), Space: O(1)
func maxProfit1(prices []int) int {
	if len(prices) < 2 {
		return 0 // 如果数组长度小于2，直接返回0
	}
	maxProfit := 0 // 初始化最大利润为0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ { // 卖出时机要在买入时机之后
			maxProfit = max(maxProfit, prices[j]-prices[i]) // 取当前利润和当前最大利润maxProfit的最大值
		}
	}
	return maxProfit // 返回最大利润
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// Time: O(n), Space: O(1)
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0 // 如果数组长度小于2，直接返回0
	}
	maxProfit := 0                     // 初始化最大利润为0
	buy := prices[0]                   // 买入价格为第0个元素的价格
	for i := 1; i < len(prices); i++ { // 接着从第一个元素开始遍历数组
		// 如果价格小于买入价格，则更新买入价格
		if prices[i] < buy {
			buy = prices[i]
		} else { // 否者计算当前最大利润
			maxProfit = max(maxProfit, prices[i]-buy)
		}
	}
	return maxProfit // 循环结束后返回最大利润即可
}
