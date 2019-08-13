package _9_Leetcode_Problems_about_Map_and_Set

func intersect(nums1 []int, nums2 []int) []int {
	var res []int
	nums1Map := make(map[int]int)

	for _, num := range nums1 {
		nums1Map[num]++
	}

	for _, num := range nums2 {
		if nums1Map[num] > 0 {
			res = append(res, num)
			nums1Map[num]--
		}
	}
	return res
}
