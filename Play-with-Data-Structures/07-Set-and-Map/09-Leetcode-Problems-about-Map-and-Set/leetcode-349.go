package _9_Leetcode_Problems_about_Map_and_Set

func intersection(nums1 []int, nums2 []int) []int {
	var res []int
	uniqueNums1 := make(map[int]bool)

	for _, num := range nums1 {
		uniqueNums1[num] = true
	}

	for _, num := range nums2 {
		if uniqueNums1[num] == true {
			res = append(res, num)
			uniqueNums1[num] = false
		}
	}

	return res
}

/**
https://leetcode-cn.com/problems/two-sum/solution/mapfang-fa-by-15821787840/
*/
func intersection1(nums1 []int, nums2 []int) []int {
	set := make(map[int]bool)
	res := make([]int, 0)
	for _, v1 := range nums1 {
		set[v1] = true
	}
	for _, v2 := range nums2 {
		if m, ok := set[v2]; ok && m { //nums2里面包含nums1里的元素
			res = append(res, v2)
			set[v2] = false //防止重复输出
		}
	}
	return res
}
