package _2_BFS

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGraphBFS(t *testing.T) {
	if g, err := NewGraph("g.txt"); err == nil {
		graphBfs := NewGraphBFS(g)
		fmt.Println("BFS Order: ", graphBfs.Order()) // BFS Order: [0 1 2 3 4 6 5]
		assert.Equal(t, graphBfs.Order(), []int{0, 1, 2, 3, 4, 6, 5})
	}
}
