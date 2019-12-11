package LeetCode

import "sort"

// Time Complexity: O(nlogn) Space Complexity: O(n)
func intersect1(nums1 []int, nums2 []int) []int {
	record := make(map[int]int)       // 记录元素和出现的频次 map的实现O(logn)
	for i := 0; i < len(nums1); i++ { // 遍历一遍数组O(n)
		record[nums1[i]]++ // 记录元素对应的频次
	}
	var resultVector []int // O(n)
	// O(nlogn)
	for i := 0; i < len(nums2); i++ {
		if record[nums2[i]] > 0 {
			resultVector = append(resultVector, nums2[i])
			record[nums2[i]]--
		}
	}
	return resultVector
}

/// Sorting and Two Pointers
/// Time Complexity: O(nlogn)
/// Space Complexity: O(1)
func intersect2(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	var res []int
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			res = append(res, nums1[i])
			i, j = i+1, j+1
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}
	return res
}
