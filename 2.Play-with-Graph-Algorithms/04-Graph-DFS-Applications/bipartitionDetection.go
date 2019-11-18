package main

import (
	"fmt"

	Graph "github.com/custergo/study_algo/graph/utils"
)

// 二分图检测: 染色过程
type BipartitionDetection struct {
	G           *Graph.AdjList
	visited     map[int]bool // 记录每个节点是否被遍历过
	colors      []int        // int型数组记录染色
	isBipartite bool         // 是否是二分图
}

func NewBipartitionDetection(G *Graph.AdjList) *BipartitionDetection {
	isBipartite := true
	colors := make([]int, G.V)
	for i := 0; i < G.V; i++ {
		colors[i] = -1 // -1 表示该顶点还未被染色
	}
	visited := make(map[int]bool, 0) // 记录每个节点是否被遍历过
	for v := 0; v < G.V; v++ {       // 遍历每个顶点，避免忽略联通分量
		if !visited[v] {
			if !dfs(v, G, visited, 0, colors) {
				isBipartite = false // 不是二分图
				break
			}
		}
	}

	return &BipartitionDetection{G: G, visited: visited, colors: colors, isBipartite: isBipartite}
}

func dfs(v int, G *Graph.AdjList, visited map[int]bool, color int, colors []int) bool {
	visited[v] = true
	colors[v] = color

	for w := G.Adjacency(v).Front(); w != nil; w = w.Next() {
		ww := w.Value.(int) // 遍历顶点v的所有相邻点
		if !visited[ww] {   // 如果没有遍历过，就递归调用
			if dfs(ww, G, visited, 1-color, colors) { // color是0传入的就是1
				return false // 如果不是二分图直接返回false
			} else if colors[ww] == colors[v] { // v的相邻顶点w已经被访问过了
				return false // 如果w顶点已经被访问过了，那么w和v的颜色相同则不是二分图
			}
		}
	}
	return true
}

func (b *BipartitionDetection) IsBipartite() bool {
	return b.isBipartite
}

func main() {
	if g, err := Graph.NewAdjList("g.txt"); err == nil {
		b := NewBipartitionDetection(g)
		fmt.Println("是否是二分图: ", b.IsBipartite()) // 是否是二分图: true
	}

	if g, err := Graph.NewAdjList("g3.txt"); err == nil {
		b := NewBipartitionDetection(g)
		fmt.Println("是否是二分图: ", b.IsBipartite()) // 是否是二分图: false
	}
}
