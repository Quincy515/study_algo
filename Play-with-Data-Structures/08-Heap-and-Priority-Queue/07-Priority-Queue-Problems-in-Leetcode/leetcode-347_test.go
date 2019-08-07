package Leetcode_347

import (
	"fmt"
	"testing"
)

func TestLeetCode(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	fmt.Println(TopKFrequent(nums, k))
}
