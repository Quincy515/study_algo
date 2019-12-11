package _7_BFS_and_Shortest_Path

import (
	"fmt"
	"testing"
)

//
func TestReadGraph(t *testing.T) {
	fmt.Println("=======测试寻路算法======")
	filename := "testG.txt"
	g := NewSparseGraph(7, false)
	readGraph := &ReadGraph{}

	if err := readGraph.Read(g, filename); err != nil {
		fmt.Println(err)
		return
	}
	g.Show()
	fmt.Println()

	fmt.Println("===比较使用深度优先遍历和广度优先遍历获得路径的不同===")
	dfs := NewPath(g, 0)
	fmt.Printf("DFS Path from 0 to %d: ", 6)
	dfs.showPath(6)

	fmt.Println("===广度优先遍历获得的是无权图的最短路径===")
	bfs := NewShortestPath(g, 0)
	fmt.Printf("BFS Path from 0 to %d: ", 6)
	bfs.showPath(6)
	fmt.Println()

	filename = "testG1.txt"
	g2 := NewSparseGraph(13, false)
	if err := readGraph.Read(g2, filename); err != nil {
		fmt.Println(err)
		return
	}
	g2.Show()
	fmt.Println()

	fmt.Println("===比较使用深度优先遍历和广度优先遍历获得路径的不同===")
	dfs2 := NewPath(g2, 0)
	fmt.Printf("DFS Path from 0 to %d: ", 3)
	dfs2.showPath(3)

	fmt.Println("===广度优先遍历获得的是无权图的最短路径===")
	bfs2 := NewShortestPath(g, 0)
	fmt.Printf("BFS Path from 0 to %d: ", 3)
	bfs2.showPath(3)
	fmt.Println()
}
