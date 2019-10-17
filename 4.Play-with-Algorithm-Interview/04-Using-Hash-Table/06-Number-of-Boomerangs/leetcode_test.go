package leetcode

import (
	"fmt"
	"testing"
)

func TestLeetCode447(t *testing.T) {
	points := [][]int{
		{0, 0},
		{1, 0},
		{2, 0},
	}
	fmt.Println(numberOfBoomerangs(points))
}

func TestLeetCode149(t *testing.T) {
	points := [][]int{{1, 1}, {2, 2}, {3, 3}}
	fmt.Println(maxPoints(points))
}
