package _1_Leetcode_Graph_Basics

import (
	"fmt"
	"testing"
)

func TestIsBipartite(t *testing.T) {
	g := [][]int{
		{1, 3},
		{0, 2},
		{1, 3},
		{0, 2},
	}
	fmt.Println(IsBipartite(g))
}
