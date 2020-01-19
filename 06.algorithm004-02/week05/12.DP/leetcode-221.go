package leetcode

// 使用84、85直方图方法 Time: O(m*n), Space: O(n)
func maximalSquareHistogram(matrix [][]byte) int {
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
		max = maximum(max, largestSquareInHistogram(heights)) // 较大值更新max
	}
	return max // 最后返回max即可
}

// 根据直方图高度计算最大正方形面积
func largestSquareInHistogram(heights []int) int {
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
		len := minimum(heights[i], r-l-1) // 矩形中宽度和高度较小值
		max = maximum(max, len*len)       // 所有固定高度的最大面积
	}
	return max
}

func maximum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minimum(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 辅助函数求三个整数的最小值
func min(a, b, c int) int {
	return minimum(minimum(a, b), c)
}

// 动态规划 Time: O(m*n), Space: O(m*n)
func maximalSquare(matrix [][]byte) int {
	// 处理边界情况，矩阵为空或长度为0，直接返回0
	if matrix == nil || len(matrix) == 0 ||
		matrix[0] == nil || len(matrix[0]) == 0 {
		return 0
	}
	// 否则把矩阵两个维度的大小分别赋值给m和n
	m, n := len(matrix), len(matrix[0])
	// 同时定义大小为m*n的二维辅助数组d
	d := make([][]int, m)
	for i := range d {
		d[i] = make([]int, n)
	}
	length := 0 // 初始化正方形边长为0

	for i := 0; i < m; i++ { // i从0遍历到m-1
		for j := 0; j < n; j++ { // j从0遍历到n-1
			if i == 0 || j == 0 || matrix[i][j] == '0' {
				if matrix[i][j] == '1' {
					d[i][j] = 1
				} else {
					d[i][j] = 0
				}
			} else { // 矩阵的字符为1或者不在第一行或第一列
				// 这时d(i,j)等于它上边，左边，左上角三个状态的最小值加一
				d[i][j] = min(d[i-1][j], d[i][j-1], d[i-1][j-1]) + 1
			}
			// 接下来用最大边长和d(i,j)进行对比
			length = maximum(length, d[i][j]) // 较大值更新边长
		}
	}
	return length * length // 循环结束后返回边长乘以边长作为正方形的面积
}

// 动态规划优化 Time: O(m*n), Space: O(n)
func maximalSquare(matrix [][]byte) int {
	// 处理边界情况，矩阵为空或长度为0，直接返回0
	if matrix == nil || len(matrix) == 0 ||
		matrix[0] == nil || len(matrix[0]) == 0 {
		return 0
	}
	// 否则把矩阵两个维度的大小分别赋值给m和n
	m, n := len(matrix), len(matrix[0])
	// 同时定义大小为n的辅助数组d
	d := make([]int, n)
	length, pre := 0, 0 // 初始化正方形边长为0

	for i := 0; i < m; i++ { // i从0遍历到m-1
		for j := 0; j < n; j++ { // j从0遍历到n-1
			tmp := d[j]
			if i == 0 || j == 0 || matrix[i][j] == '0' {
				if matrix[i][j] == '1' {
					d[j] = 1
				} else {
					d[j] = 0
				}
			} else { // 矩阵的字符为1或者不在第一行或第一列
				// 这时d(i,j)等于它上边，左边，左上角三个状态的最小值加一
				d[j] = min(pre, d[j], d[j-1]) + 1
			}
			// 接下来用最大边长和d(i,j)进行对比
			length = maximum(length, d[j]) // 较大值更新边长
			pre = tmp
		}
	}
	return length * length // 循环结束后返回边长乘以边长作为正方形的面积
}
