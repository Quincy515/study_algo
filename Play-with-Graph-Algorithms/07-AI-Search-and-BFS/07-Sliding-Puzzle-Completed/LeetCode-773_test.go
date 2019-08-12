package LeetCode_773

import (
	"fmt"
	"testing"
)

func TestLeetCode(t *testing.T) {
	board := [][]int{{1, 2, 3}, {4, 0, 5}}
	fmt.Println(slidingPuzzle(board)) // 1

	board1 := [][]int{{1, 2, 3}, {5, 4, 0}}
	fmt.Println(slidingPuzzle(board1)) // -1

	board2 := [][]int{{4, 1, 2}, {5, 0, 3}}
	fmt.Println(slidingPuzzle(board2)) // 5

	board3 := [][]int{{3, 2, 4}, {1, 5, 0}}
	fmt.Println(slidingPuzzle(board3)) // 14
}
