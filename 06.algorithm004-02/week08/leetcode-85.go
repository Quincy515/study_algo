package leetcode

// Time: O(m*n), Space: O(n)
func maximalRectangle(matrix [][]byte) int {
	if matrix == nil || len(matrix) == 0 || matrix[0] == nil || len(matrix[0]) == 0 {
		return 0 // 矩阵为空或长度为0
	}
	m, n := len(matrix), len(matrix[0]) // 把矩阵两个维度的大小分别赋值给m和n
	heights := make([]int, n)           // 同时定义大小为n的高度数组
	max := 0                            // 定义最大矩形面积为0
	for i := 0; i < m; i++ {            // 遍历矩阵第一个维度,选取一行作为直方图的底边
		for j := 0; j < n; j++ { // 遍历这一行上的n个字符
			if matrix[i][j] == '1' { // 如果字符为1
				heights[j] = heights[j] + 1 // 高度值在原来基础上累积1
			} else { // 如果字符为0
				heights[j] = 0 // 高度值重置为0
			}
		}
		// 调用辅助函数,计算当前高度数组产生的最大矩形面积,并和max进行对比
		max = maximum(max, largestRectangleArea(heights)) // 较大值更新max
	}
	return max // 最后返回max即可
}

// 根据直方图高度计算最大矩形面积
func largestRectangleArea(heights []int) int {
	if heights == nil || len(heights) == 0 {
		return 0
	}
	max, n := 0, len(heights)
	for i := 0; i < n; i++ { // 对每个固定高度求其最大面积
		l, r := i, i // 固定高度，扩展左右边界，求其最大面积的宽度
		for l >= 0 && heights[l] >= heights[i] {
			l-- // 左边界跑出边界或者小于固定高度值为止
		}
		for r < n && heights[r] >= heights[i] {
			r++ // 右边界跑出边界或者小于固定高度值为止
		}
		max = maximum(max, heights[i]*(r-l-1)) // 所有固定高度的最大面积
	}
	return max
}

func maximum(a, b int) int {
	if a > b {
		return a
	}
	return b
}
