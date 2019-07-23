package main

import (
	"fmt"

	Graph "github.com/custergo/study_algo/graph/utils"
)

func GraphDFS(G *Graph.AdjList) []int {
	var order []int                  // 存放遍历结果
	visited := make(map[int]bool, 0) // 记录每个节点是否被遍历过
	dfs(0, G, &order, visited)
	return order
}

func dfs(v int, G *Graph.AdjList, order *[]int, visited map[int]bool) {
	visited[v] = true
	*order = append(*order, v)

	for w := G.Adjacency(v).Front(); w != nil; w = w.Next() {
		ww := w.Value.(int)
		if !visited[ww] {
			dfs(ww, G, order, visited)
		}
	}
}

func main() {
	if g, err := Graph.NewAdjList("g.txt"); err == nil {
		fmt.Println(GraphDFS(g))
	}
}
