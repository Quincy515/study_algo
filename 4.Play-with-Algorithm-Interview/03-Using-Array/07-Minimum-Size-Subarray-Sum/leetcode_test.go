package LeetCode

import (
	"fmt"
	"testing"
)

func TestLeetCode209(t *testing.T) {
	s, nums := 7, []int{2, 3, 1, 2, 4, 3}
	fmt.Println(minSubArrayLen1(s, nums))
	s1, nums1 := 15, []int{1, 2, 3, 4, 5}
	fmt.Println(minSubArrayLen2(s1, nums1))
	fmt.Println(minSubArrayLen3(s1, nums1))
	fmt.Println(minSubArrayLen4(s1, nums1))
	fmt.Println(minSubArrayLen5(s1, nums1))
	fmt.Println(minSubArrayLen6(s1, nums1))
}
