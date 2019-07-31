package _3_Recursion_Basics

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(Sum(nums)) // 36
}
