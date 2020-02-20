package leetcode

func containsDuplicate(nums []int) bool {
	record := make(map[int]int) // hash 表记录出现过的元素
	for i := 0; i < len(nums); i++ {
		if _, ok := record[nums[i]]; ok {
			return true
		}
		record[nums[i]] = i
	}
	return false
}
