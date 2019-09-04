package LeetCode

import (
	"fmt"
	"testing"
)

func TestSortColors(t *testing.T) {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors1(nums)
	fmt.Println(nums)
}

func TestMerge(t *testing.T) {
	nums := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3

	nums1 := make([]int, m+n)
	nums4 := make([]int, m+n)
	nums3 := make([]int, m+n)

	copy(nums1, nums)
	copy(nums3, nums)
	copy(nums4, nums)

	merge1(nums1, m, nums2, n)
	fmt.Println(nums1)
	merge2(nums3, m, nums2, n)
	fmt.Println(nums3)
	merge3(nums4, m, nums2, n)
	fmt.Println(nums4)
}

func TestFindKthLargest(t *testing.T) {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	fmt.Println(findKthLargest1(nums, k))
	fmt.Println(findKthLargest2(nums, k))
	fmt.Println(findKthLargest3(nums, k))
}
