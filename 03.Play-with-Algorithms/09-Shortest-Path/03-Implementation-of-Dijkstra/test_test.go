package Dijkstra

import (
	"fmt"
	"testing"
)

func TestGraph(t *testing.T) {
	fmt.Println("===读取 testG1.txt 生成 Weighted Sparse Graph===")
	v, filename := 5, "testG1.txt"

	g1 := NewSparseGraph(v, true) // Dijkstra算法同样适用于有向图
	readGraph := &ReadGraph{}
	if err := readGraph.Read(g1, filename); err != nil {
		fmt.Println(err)
		return
	}
	g1.Show()
	fmt.Println()

	// Test Dijkstra
	fmt.Println("Test Dijkstra: ")
	dij := NewDijkstra(g1, 0)
	for i := 1; i < v; i++ {
		if dij.hashPathTo(i) {
			fmt.Println("Shortes Path to ", i, " : ", dij.shortestPathTo(i))
			dij.showPath(i)
		} else {
			fmt.Println("No Path to ", i)
		}
		fmt.Println("--------------")
	}
}

/*
=== RUN   TestGraph
===读取 testG1.txt 生成 Weighted Sparse Graph===
vertex 0:	( to:1, wt: 5)	( to:2, wt: 2)	( to:3, wt: 6)
vertex 1:	( to:0, wt: 5)	( to:4, wt: 1)	( to:2, wt: 1)
vertex 2:	( to:0, wt: 2)	( to:1, wt: 1)	( to:4, wt: 5)	( to:3, wt: 3)
vertex 3:	( to:0, wt: 6)	( to:2, wt: 3)	( to:4, wt: 2)
vertex 4:	( to:1, wt: 1)	( to:2, wt: 5)	( to:3, wt: 2)

Test Dijkstra:
Shortes Path to  1  :  3
0 -> 2 -> 1
Shortes Path to  2  :  2
0 -> 2
Shortes Path to  3  :  5
0 -> 2 -> 3
Shortes Path to  4  :  4
0 -> 2 -> 1 -> 4
*/
