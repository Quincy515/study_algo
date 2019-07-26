package _3_Single_Source_Path_BFS

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGraphBFS(t *testing.T) {
	if g, err := NewGraph("g.txt"); err == nil {
		sspath := NewSingleSourcePath(g, 0)     // 图g和源点s
		fmt.Println("0 -> 6: ", sspath.Path(6)) // 0 -> 6: [0 2 6]
		assert.Equal(t, sspath.Path(6), []int{0, 2, 6})
	}
}
