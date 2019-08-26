package _6_Finding_a_Path

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

	p := NewPath(g, 0)
	fmt.Printf("Path from 0 to %d: ", 6)
	p.showPath(6)
}
