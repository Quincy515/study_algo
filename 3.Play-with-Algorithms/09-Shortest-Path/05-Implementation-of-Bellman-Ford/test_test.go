package Bellman_Ford

import (
	"fmt"
	"testing"
)

func TestGraph(t *testing.T) {
	fmt.Println("===读取 testG2.txt 生成 Weighted Sparse Graph===")
	v, s, filename := 5, 0, "testG2.txt"
	//v, s, filename := 5, 0, "testG_negative_circle.txt" // The graph contain negative cycle!

	g1 := NewSparseGraph(v, true)
	readGraph := &ReadGraph{}
	if err := readGraph.Read(g1, filename); err != nil {
		fmt.Println(err)
		return
	}
	//g1.Show()
	fmt.Println()

	// Test Bellman-Ford算法
	fmt.Println("测试Bellman-Ford算法: ")
	bellmanFord := NewBellmanFord(g1, s)
	if bellmanFord.negativeCycle() {
		fmt.Println("The graph contain negative cycle!")
	} else {
		for i := 0; i < v; i++ {
			if i == s {
				continue
			}
			if bellmanFord.hasPathTo(i) {
				fmt.Println("Shortest Path to ", i, " : ", bellmanFord.shortestPathTo(i))
				bellmanFord.showPath(i)
			} else {
				fmt.Println("No Path to ", i)
			}
			fmt.Println("----------------")
		}
	}
}

/*
=== RUN   TestGraph
===读取 testG2.txt 生成 Weighted Sparse Graph===
vertex 0:	( to:1, wt: 5)	( to:2, wt: 2)	( to:3, wt: 6)
vertex 1:	( to:2, wt: -4)	( to:4, wt: 2)
vertex 2:	( to:4, wt: 5)	( to:3, wt: 3)
vertex 3:
vertex 4:	( to:3, wt: -3)

测试Bellman-Ford算法:
Shortest Path to  1  :  5
0 -> 1
----------------
Shortest Path to  2  :  1
0 -> 1 -> 2
----------------
Shortest Path to  3  :  3
0 -> 1 -> 2 -> 4 -> 3
----------------
Shortest Path to  4  :  6
0 -> 1 -> 2 -> 4
----------------
--- PASS: TestGraph (0.00s)
PASS
*/
