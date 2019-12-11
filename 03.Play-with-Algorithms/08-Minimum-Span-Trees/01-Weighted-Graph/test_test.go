package _1_Weighted_Graph

import (
	"fmt"
	"testing"
)

func TestGraph(t *testing.T) {
	fmt.Println("===Test Weighted Dense Graph===")
	filename := "testG1.txt"
	g1 := NewDenseGraph(8, false)
	readGraph := &ReadGraph{}
	if err := readGraph.Read(g1, filename); err != nil {
		fmt.Println(err)
		return
	}
	g1.Show()
	fmt.Println()

	fmt.Println("===Test Weighted Sparse Graph===")
	g2 := NewSparseGraph(8, false)
	if err := readGraph.Read(g2, filename); err != nil {
		fmt.Println(err)
		return
	}
	g2.Show()
	fmt.Println()
}
