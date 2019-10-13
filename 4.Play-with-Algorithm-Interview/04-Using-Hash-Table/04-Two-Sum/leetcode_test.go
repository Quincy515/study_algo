package LeetCode

import (
	"fmt"
	"testing"
)

func TestLeetCode1(t *testing.T) {
	nums1, target1 := []int{2, 7, 11, 15}, 9
	fmt.Println(twoSum1(nums1, target1))
	nums2, target2 := []int{3, 2, 4}, 6
	fmt.Println(twoSum1(nums2, target2))
}

func TestLeetCode15(t *testing.T) {
	nums1 := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum1(nums1))
	nums2 := []int{0, 0, 0, 0}
	fmt.Println(threeSum1(nums2))
	nums3 := []int{-1, 0, 1, 0}
	fmt.Println(threeSum1(nums3))
}

func TestLeetCode16(t *testing.T) {
	nums1, target1 := []int{-1, 2, 1, -4}, 1
	nums2, target2 := []int{1, 1, 1, 0}, 100
	fmt.Println(threeSumClosest1(nums1, target1))
	fmt.Println(threeSumClosest2(nums1, target1))
	fmt.Println(threeSumClosest3(nums1, target1))
	fmt.Println(threeSumClosest4(nums2, target2))
}

func TestLeetCode18(t *testing.T) {
	nums, target := []int{1, 0, -1, 0, -2, 2}, 0
	fmt.Println(fourSum1(nums, target))
	fmt.Println(fourSum2(nums, target))
}
