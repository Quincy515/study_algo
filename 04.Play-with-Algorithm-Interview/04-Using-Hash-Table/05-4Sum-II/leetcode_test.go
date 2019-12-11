package LeetCode

import (
	"fmt"
	"testing"
)

func TestLeetCode454(t *testing.T) {
	A, B, C, D := []int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}
	fmt.Println(fourSumCount2(A, B, C, D))
	A2, B2, C2, D2 := []int{0}, []int{0}, []int{0}, []int{0}
	fmt.Println(fourSumCount2(A2, B2, C2, D2))
	A3, B3, C3, D3 := []int{-1, -1}, []int{-1, 1}, []int{-1, 1}, []int{1, -1}
	fmt.Println(fourSumCount2(A3, B3, C3, D3))
}

func TestLeetCode49(t *testing.T) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams2(strs))
}
