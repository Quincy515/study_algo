package LeetCode

import (
	"fmt"
	"testing"
)

func TestLeetCode349(t *testing.T) {
	nums1, nums2 := []int{1, 2, 2, 1}, []int{2, 2}
	fmt.Println(intersection1(nums1, nums2))
	fmt.Println(intersection2(nums1, nums2))
}
