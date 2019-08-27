package Lazy_Prim

import (
	"fmt"
	"testing"
)

func TestGraph(t *testing.T) {
	fmt.Println("===读取 testG1.txt 生成 Weighted Sparse Graph===")
	filename := "testG1.txt"
	g1 := NewSparseGraph(8, false)
	readGraph := &ReadGraph{}
	if err := readGraph.Read(g1, filename); err != nil {
		fmt.Println(err)
		return
	}
	g1.Show()
	fmt.Println()

	// Test Lazy Prim MST
	fmt.Println("Test Lazy Prim MST: ")
	p := NewLazyPrimMST(g1)
	mst := p.mstEdges()
	for i := 0; i < len(mst); i++ {
		fmt.Println(mst[i])
	}
	fmt.Println("The MST weight is: ", p.result())
}
