package main

import (
	"fmt"

	Graph "github.com/custergo/study_algo/graph/utils"
)

func CC(G *Graph.AdjList) int {
	cccount := 0                     // 联通分量的个数
	visited := make(map[int]bool, 0) // 记录每个节点是否被遍历过
	for v := 0; v < G.V; v++ {       // 遍历每个顶点，避免忽略联通分量
		if !visited[v] {
			dfs(v, G, visited)
			cccount++
		}
	}
	return cccount
}

func dfs(v int, G *Graph.AdjList, visited map[int]bool) {
	visited[v] = true

	for w := G.Adjacency(v).Front(); w != nil; w = w.Next() {
		ww := w.Value.(int) // 遍历顶点v的所有相邻点
		if !visited[ww] {   // 如果没有遍历过，就递归调用
			dfs(ww, G, visited)
		}
	}

}

func main() {
	if g, err := Graph.NewAdjList("g.txt"); err == nil {
		cc := CC(g)
		fmt.Println("联通分量的个数: ", cc) // 联通分量的个数: [0 1 3 2 6 4 5]
	}
}
