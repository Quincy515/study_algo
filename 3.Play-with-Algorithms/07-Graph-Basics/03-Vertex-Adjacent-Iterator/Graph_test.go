package Graph_Representation

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestSparseGraph(t *testing.T) {
	n, m := 20, 100 // 定义20个节点100条边的图
	rand.Seed(time.Now().Unix())

	// Sparse Graph
	fmt.Println("----------Sparse Graph----------")
	g1 := NewSparseGraph(n, false)
	for i := 0; i < m; i++ {
		a := rand.Int() % n
		b := rand.Int() % n
		g1.AddEdge(a, b)
	}

	// 应用稀疏图的迭代器,依次遍历了每个点相应的邻边
	for v := 0; v < n; v++ {
		fmt.Print(v, " : ")
		adj := NewAdjIteratorSparseGraph(g1, v)
		for w := adj.Begin(); !adj.End(); w = adj.Next() {
			fmt.Print(w, " ")
		}
		fmt.Println()
	} // O(E)有多少边,就遍历多少次

	// Dense Graph
	fmt.Println("----------Dense Graph----------")
	g2 := NewDenseGraph(n, false)
	for i := 0; i < m; i++ {
		a := rand.Int() % n
		b := rand.Int() % n
		g2.AddEdge(a, b)
	}

	// O(v^2)
	for v := 0; v < n; v++ {
		fmt.Print(v, " : ")
		adj := NewAdjIteratorDenseGraph(g2, v)
		for w := adj.Begin(); !adj.End(); w = adj.Next() {
			fmt.Print(w, " ")
		}
		fmt.Println()
	}
}
