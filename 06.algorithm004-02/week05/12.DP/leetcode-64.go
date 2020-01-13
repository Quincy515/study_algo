package leetcode

// 二维数组表示状态的版本 Time: O(m*n), Space: O(m*n)
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0]) // 初始化矩阵大小m和n
	d := make([][]int, m)           // 定义二维数组d
	for i := range d {
		d[i] = make([]int, n)
	}
	d[0][0] = grid[0][0]     // 初始化d(0,0)
	for i := 1; i < m; i++ { // 初始化第一列
		// 每个状态值d(i,0)等于上边的状态值d(i-1,0)加上当前的矩阵元素值grid(i,0)
		d[i][0] = d[i-1][0] + grid[i][0]
	}
	for j := 1; j < n; j++ { // 同理初始化第一行
		// 每个状态值d(0,j)等于上边的状态值d(0,j-1)加上当前的矩阵元素值grid(0,j)
		d[0][j] = d[0][j-1] + grid[0][j]
	}
	// 接着从下标1开始填充状态数组
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// 状态d(i,j)是它上边状态d(i-1,j)和左边状态d(i,j-1)的较小值
			// 再加上当前矩阵元素grid[i][j]
			d[i][j] = min(d[i-1][j], d[i][j-1]) + grid[i][j]
		}
	}
	return d[m-1][n-1] // 最后返回状态数组最后的元素即可
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 只使用一位数组来表示状态的版本 Time: O(m*n), Space: O(n)
func minPathSumOn(grid [][]int) int {
	m, n := len(grid), len(grid[0]) // 初始化矩阵大小m和n
	d := make([]int, n)             // 定义一维数组d来表示状态
	d[0] = grid[0][0]               // 初始化d(0,0)
	for j := 1; j < n; j++ {        // 初始化第一行
		// 每个状态值d(0,j)等于上边的状态值d(0,j-1)加上当前的矩阵元素值grid(0,j)
		d[j] = d[j-1] + grid[0][j]
	}
	for i := 1; i < m; i++ { // 从1开始遍历每一行
		// 遍历到新的一行时,
		// d[0]需要加上这一行的第0个元素
		// 表示从上边走到了当前行的行首，第0个元素
		d[0] += grid[i][0]
		// 开始遍历当前行
		for j := 1; j < n; j++ {
			// 更新状态d[j]为d[j]和d[j-1]中的较小值再加上grid[i][j]
			d[j] = min(d[j], d[j-1]) + grid[i][j]
		}
	}
	return d[n-1] // 最后返回d[n-1]即可
}
