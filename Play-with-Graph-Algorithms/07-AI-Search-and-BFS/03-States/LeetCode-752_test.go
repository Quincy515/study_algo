package LeetCode_752

import (
	"fmt"
	"testing"
)

func TestLeetCode(t *testing.T) {
	deadends := []string{"0201", "0101", "0102", "1212", "2002"}
	target := "0202"
	fmt.Println(openLock(deadends, target))

	//deadends1 := []string{"8888"}
	//target1 := "0009"
	//fmt.Println(openLock(deadends1, target1))
}
