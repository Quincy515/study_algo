package leetcode

// 二维辅助数组的版本 Time: O(m*n), Space: O(m*n)
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0]) // 把行数和列数分别赋值给m和n
	d := make([][]int, m)                           // 定义m*n的状态数组
	for i := range d {
		d[i] = make([]int, n)
	}
	d[0][0] = 1 // 根据是否有障碍物初始化为0或1
	if obstacleGrid[0][0] == 1 {
		d[0][0] = 0
	}
	// 初始化最左侧那一列的状态
	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			d[i][0] = 0 // 如果位置上有障碍物，到达它的路径数量为0
		} else { // 否则等于上边位置的路径数量
			d[i][0] = d[i-1][0]
		}
	}
	// 初始化最上面那一行的状态
	for j := 1; j < n; j++ {
		if obstacleGrid[0][j] == 1 {
			d[0][j] = 0
		} else {
			d[0][j] = d[0][j-1]
		}
	}
	// 接着使用两层循环遍历矩阵上的其他位置
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				d[i][j] = 0 // 对于其他位置如果该位置有障碍物，它的路径数量为0
			} else {
				d[i][j] = d[i-1][j] + d[i][j-1]
			}
		}
	}
	return d[m-1][n-1]
}

// Time: O(m*n), Space: O(n)
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0]) // 把行数和列数分别赋值给m和n
	d := make([]int, n)
	if obstacleGrid[0][0] == 1 {
		d[0] = 0
	} else {
		d[0] = 1
	}

	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			d[0] = 0
		} else {
			d[0] = d[0]
		}
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				d[j] = 0
			} else {
				d[j] = d[j] + d[j-1]
			}
		}
	}
	return d[n-1]
}
