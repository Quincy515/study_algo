package _3_Graph_Construction

import (
	"container/list"
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

	// 解决最大岛屿问题-转换成图论问题
	G := constructGraph(R, C, grid)

	res := 0
	visited := make(map[int]bool, len(G))

	for v := 0; v < len(G); v++ {
		x, y := v/C, v%C
		if grid[x][y] == 1 && !visited[v] {
			res = int(math.Max(float64(res), float64(dfs(v, visited, G))))
		}
	}
	return res
}

func dfs(v int, visited map[int]bool, G []*list.List) int {
	visited[v] = true
	res := 1
	for w := G[v].Front(); w != nil; w = w.Next() {
		p := w.Value.(int)
		if !visited[p] {
			res += dfs(p, visited, G)
		}
	}
	return res
}

func constructGraph(R, C int, grid [][]int) []*list.List {
	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	g := make([]*list.List, 0) // 是一个链表数组
	for i := 0; i < R*C; i++ {
		g = append(g, list.New()) // 第i个元素表示的是第i个顶点和它相邻的顶点
	}

	for v := 0; v < len(g); v++ {
		x, y := v/C, v%C
		if grid[x][y] == 1 { // 1表示是陆地
			for d := 0; d < 4; d++ { // 4联通
				nextx, nexty := x+dirs[d][0], y+dirs[d][1]
				if inArea(nextx, nexty, R, C) && grid[nextx][nexty] == 1 {
					next := nextx*C + nexty
					g[v].PushBack(next)
					g[next].PushBack(v)
				}
			}
		}
	}
	return g
}

func inArea(x, y, R, C int) bool {
	return x >= 0 && x < R && y >= 0 && y < C
}
