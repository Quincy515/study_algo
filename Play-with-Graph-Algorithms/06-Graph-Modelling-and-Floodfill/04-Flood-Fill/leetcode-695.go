package _4_Flood_Fill

import (
	"math"
)

func MaxAreaOfIsland(grid [][]int) int {
	// 不显示的创建图来解决

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
	visited := make([][]bool, 0)
	for range make([]int, R) { // 初始化空的 visited 二维数组
		visited = append(visited, make([]bool, C))
	}

	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			if !visited[i][j] && grid[i][j] == 1 {
				res = int(math.Max(float64(res), float64(dfs(i, j, R, C, visited, grid))))
			}
		}

	}
	return res
}

func dfs(x, y, R, C int, visited [][]bool, grid [][]int) int {
	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	visited[x][y] = true
	res := 1
	for d := 0; d < 4; d++ {
		nextx, nexty := x+dirs[d][0], y+dirs[d][1]
		// 1. 合法 2. 未被遍历过 3. 是陆地
		if inArea(nextx, nexty, R, C) && !visited[nextx][nexty] && grid[nextx][nexty] == 1 {
			res += dfs(nextx, nexty, R, C, visited, grid)
		}
	}
	return res
}

func inArea(x, y, R, C int) bool {
	return x >= 0 && x < R && y >= 0 && y < C
}
