package _7_Bipartition_Detection_BFS

import (
	"fmt"
	"testing"
)

func TestNewBipartitionDetection(t *testing.T) {
	if g, err := NewGraph("g.txt"); err == nil {
		b := NewBipartitionDetection(g)
		fmt.Println("是否有环: ", b.IsBipartite()) // true
	}

	if g, err := NewGraph("g2.txt"); err == nil {
		b := NewBipartitionDetection(g)
		fmt.Println("是否有环: ", b.IsBipartite()) // false
	}

	if g, err := NewGraph("g3.txt"); err == nil {
		b := NewBipartitionDetection(g)
		fmt.Println("是否有环: ", b.IsBipartite()) // true
	}
}
