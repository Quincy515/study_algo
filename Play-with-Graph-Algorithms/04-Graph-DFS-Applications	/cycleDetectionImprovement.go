package main

import (
	"fmt"

	Graph "github.com/custergo/study_algo/graph/utils"
)

// 环的检测代码优化
type CycleDetection struct {
	G        *Graph.AdjList
	visited  map[int]bool
	hasCycle bool
}

func NewCycleDetection(G *Graph.AdjList) *CycleDetection {
	hasCycle := false                // 记录是否有环
	visited := make(map[int]bool, 0) // 记录每个节点是否被遍历过
	for v := 0; v < G.V; v++ {       // 遍历每个顶点，避免忽略联通分量
		if !visited[v] {
			if dfs(v, v, G, visited) {
				hasCycle = true
				break
			}
		}
	}
	return &CycleDetection{
		G:        G,
		visited:  visited,
		hasCycle: hasCycle,
	}
}

func dfs(v, parent int, G *Graph.AdjList, visited map[int]bool) bool {
	visited[v] = true
	for w := G.Adjacency(v).Front(); w != nil; w = w.Next() {
		ww := w.Value.(int) // 遍历顶点v的所有相邻点
		if !visited[ww] {   // 如果没有遍历过，就递归调用
			if dfs(ww, v, G, visited) {
				return true
			}
		} else if ww != parent { // 如果w被访问过了，w这个节点不能是v的parent上一个节点
			return true
		}
	}
	return false
}

func (c *CycleDetection) HasCycle() bool {
	return c.hasCycle
}

func main() {
	if g, err := Graph.NewAdjList("g.txt"); err == nil {
		cycle := NewCycleDetection(g)
		fmt.Println("图中是否有环: ", cycle.HasCycle()) // 图中是否有环: true
	}

	if g, err := Graph.NewAdjList("g2.txt"); err == nil {
		cycle := NewCycleDetection(g)
		fmt.Println("图中是否有环: ", cycle.HasCycle()) // 图中是否有环: false
	}
}
