package leetcode

// 滑动窗口+查找表 Time: O(n), Space: O(k)
func containsNearbyDuplicate(nums []int, k int) bool {
	record := make(map[int]int) // 记录窗口的查找表
	for i := 0; i < len(nums); i++ {
		if _, ok := record[nums[i]]; ok {
			return true // 在窗口中找到这个元素
		}
		// 否则说明新的数与窗口中任意数都不重复
		record[nums[i]] = i // 将新的数添加到窗口中
		// 判断这个窗口有多大,保持record中最多有k个元素
		if len(record) == k+1 {
			// 删除这个窗口最左侧的数据
			delete(record, nums[i-k])
		}
	}
	return false // 循环结束后没有返回true即没有找到满足条件
}
