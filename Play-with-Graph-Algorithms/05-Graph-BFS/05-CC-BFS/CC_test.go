package _5_CC_BFS

import (
	"fmt"
	"testing"
)

func TestCC(t *testing.T) {
	if g, err := NewGraph("g.txt"); err == nil {
		cc := NewCC(g)
		fmt.Println(cc.Count()) // 2

		fmt.Println(cc.isConnected(0, 6)) // true
		fmt.Println(cc.isConnected(5, 6)) // false

		comp := cc.components()
		for ccid := 0; ccid < len(comp); ccid++ {
			fmt.Print(ccid, " : ")         // 0 : 0 1 2 3 4 6
			for _, w := range comp[ccid] { // 1 : 5
				fmt.Print(w, " ")
			}
			fmt.Println()
		}

	}
}
