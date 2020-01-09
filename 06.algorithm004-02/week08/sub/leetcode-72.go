package leetcode

import "math"

// Time: O(m*n) Space: O(m*n)
func minDistance(word1 string, word2 string) int {
	s, t := word1, word2
	m, n := len(s)+1, len(t)+1 // 定义二维数组长和宽
	d := make([][]int, m)      // 定义二维辅助数组d
	for i := range d {
		d[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		d[i][0] = i // 行
	}
	for j := 0; j < n; j++ {
		d[0][j] = j // 列
	}

	// 二维数组的填充
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if s[i-1] == t[j-1] {
				d[i][j] = d[i-1][j-1]
			} else {
				d[i][j] = min(d[i-1][j-1], d[i-1][j], d[i][j-1]) + 1
			}
		}
	}
	return d[m-1][n-1]
}

func min(a, b, c int) int {
	return int(math.Min(float64(a), math.Min(float64(b), float64(c))))
}
