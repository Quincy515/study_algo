package Leetcode_1091

import (
	"fmt"
	"testing"
)

func TestLeetCode(t *testing.T) {
	grid := [][]int{{0, 1}, {1, 0}}
	fmt.Println(shortestPathBinaryMatrix(grid))

	grid1 := [][]int{
		{0, 0, 0},
		{1, 1, 0},
		{1, 1, 0},
	}
	fmt.Println(shortestPathBinaryMatrix(grid1))

	grid2 := [][]int{
		{0, 1, 0, 0, 0},
		{0, 1, 0, 0, 0},
		{0, 0, 0, 0, 1},
		{0, 1, 1, 1, 0},
		{0, 1, 0, 0, 0},
	}
	fmt.Println(shortestPathBinaryMatrix(grid2))

}
