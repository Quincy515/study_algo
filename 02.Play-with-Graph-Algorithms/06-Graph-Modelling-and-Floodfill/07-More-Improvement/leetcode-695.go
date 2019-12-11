package _7_More_Improvement

import (
	"math"
)

func MaxAreaOfIsland(grid [][]int) int {
	if grid == nil {
		return 0
	}
	R := len(grid) // 行数
	if R == 0 {
		return 0
	}

	C := len(grid[0]) // 列数
	if C == 0 {
		return 0
	}

	res := 0

	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			if grid[i][j] == 1 {
				res = int(math.Max(float64(res), float64(dfs(i, j, R, C, grid))))
			}
		}

	}
	return res
}

func dfs(x, y, R, C int, grid [][]int) int {
	if x < 0 || x >= R || y < 0 || y >= C || grid[x][y] == 0 {
		return 0
	}
	grid[x][y] = 0
	res := 1
	res += dfs(x-1, y, R, C, grid)
	res += dfs(x, y+1, R, C, grid)
	res += dfs(x+1, y, R, C, grid)
	res += dfs(x, y-1, R, C, grid)
	return res
}
