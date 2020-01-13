package leetcode

// 动态规划 Time: O(m*n), Space: O(m*n)
func uniquePaths(m int, n int) int {
	if m < 1 || n < 1 {
		return 0 // 如果m和n小于1，则没有路径可走，返回1
	}
	d := make([][]int, m) // 定义状态数组d
	for i := range d {
		d[i] = make([]int, n)
	}
	// 初始化第0列和第0行路径数量都为1
	for i := 0; i < m; i++ {
		d[i][0] = 1
	}
	for j := 0; j < n; j++ {
		d[0][j] = 1
	}

	// 使用两层循环计算所有其他状态d(i,j)
	// 为上边和左边的路径之和
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			d[i][j] = d[i-1][j] + d[i][j-1]
		}
	}
	return d[m-1][n-1] // 最后返回二维数组右下角的值即可
}

// 组合数学方法 Time: O(min(m, n)), Space: O(1)
func uniquePathsMath(m int, n int) int {
	if m < 1 || n < 1 {
		return 0 // 如果m和n小于1，则没有路径可走，返回1
	}
	// 在向下的m-1步和向右的n-1步选取较小值
	small := min(m-1, n-1)
	total := m + n - 2 // 定义总共的步数是m+n-2
	result := 1        // 定义最后的结果result
	// 循环small次来计算结果，每次计算都乘以分子中的一项，然后除以分母中的一项
	for i := 0; i < small; i++ {
		result = result * (total - i) / (i + 1)
	}
	return result // 最后返回结果即可
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
