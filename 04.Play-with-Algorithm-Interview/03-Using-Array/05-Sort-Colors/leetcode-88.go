package LeetCode

import (
	"fmt"
	"sort"
)

// 先合并再排序 Time Complexity: O(nlogn) Space Complexity: O(1)
func merge1(nums1 []int, m int, nums2 []int, n int) {
	nums1 = append(nums1[:m], nums2...)
	sort.Ints(nums1)
}

func merge2(nums1 []int, m int, nums2 []int, n int) {
	if len(nums1) != m+n || len(nums2) != n {
		panic("断言失败")
	}
	for i := 0; i < n; i++ {
		nums1[m+i] = nums2[i]
	}
	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] < nums1[j]
	})
}

// Standard Merge process in merge sort
// Time Complexity: O(n) Space Complexity: O(1)
func merge3(nums1 []int, m int, nums2 []int, n int) {
	for i := n + m - 1; i >= n; i-- {
		nums1[i] = nums1[i-n]
	}
	fmt.Println(nums1)

	i := n // pointer for nums1 [n,n+m)
	j := 0 // pointer for nums2 [0,n)
	k := 0 // pointer merged nums1 [0,n+m)
	for k < n+m {
		if i >= n+m {
			nums1[k] = nums2[j]
			k, j = k+1, j+1
		} else if j >= n {
			nums1[k] = nums1[i]
			k, i = k+1, i+1
		} else if nums1[i] < nums2[j] {
			nums1[k] = nums1[i]
			k, i = k+1, i+1
		} else {
			nums1[k] = nums2[j]
			k, j = k+1, j+1
		}
	}
}
