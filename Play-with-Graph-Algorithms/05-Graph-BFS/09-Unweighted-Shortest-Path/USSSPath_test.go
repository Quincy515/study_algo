package _9_Unweighted_Shortest_Path

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGraphBFS(t *testing.T) {
	if g, err := NewGraph("g.txt"); err == nil {
		ussspath := NewUSSSPath(g, 0)             // 图g和源点s
		fmt.Println("0 -> 6: ", ussspath.Path(6)) // 0 -> 6: [0 2 6]
		assert.Equal(t, ussspath.Path(6), []int{0, 2, 6})
		fmt.Println(ussspath.Dis(6)) // 0 -> 6: 需要经过两条边
	}
}
