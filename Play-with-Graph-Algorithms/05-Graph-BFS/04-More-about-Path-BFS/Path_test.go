package _4_More_about_Path_BFS

import (
	"fmt"
	"testing"
)

func TestNewGraphBFS(t *testing.T) {
	if g, err := NewGraph("g.txt"); err == nil {
		path := NewPath(g, 0, 6)             // 图g和源点s
		fmt.Println("0 -> 6: ", path.Path()) // 0 -> 1: [0 1]
	}
}
