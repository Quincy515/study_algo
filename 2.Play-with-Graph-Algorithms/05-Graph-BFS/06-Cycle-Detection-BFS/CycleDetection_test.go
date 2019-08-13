package _6_Cycle_Detection_BFS

import (
	"fmt"
	"testing"
)

func TestNewCycleDetection(t *testing.T) {
	if g, err := NewGraph("g.txt"); err == nil {
		cycle := NewCycleDetection(g)
		fmt.Println("是否有环: ", cycle.HasCycle()) // true
	}
	if g, err := NewGraph("g2.txt"); err == nil {
		cycle := NewCycleDetection(g)
		fmt.Println("是否有环: ", cycle.HasCycle()) // false
	}
}
