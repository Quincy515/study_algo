package leetcode

// Time Complexity: O(n^2), Space Complexity: O(n)
func lengthOfLIS(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	m := 1 // 定义最大长度为1
	// d[i]表示以第i个数结尾的最长上升子序列的长度
	d := make([]int, len(nums))
	d[0] = 1 // 初始化d[0],下标为0的数的最长递增子序列长度为1

	for i := 1; i < len(nums); i++ { // 遍历整个数组
		for j := 0; j < i; j++ { // 遍历第i个数字前面的所有数
			cur := 1
			if nums[i] > nums[j] { // nums[i]与i前面的数nums[j]一一对比
				cur = d[j] + 1 // 如果nums[i]比较大，则候选长度+1
			}
			d[i] = max(d[i], cur) // 取所有候选长度的最大值为第i个数字的最长递增长度
		}
		m = max(m, d[i]) // 更新最长递增子序列的长度
	}
	return m
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Time: O(n*log(n)), Space: O(n)
func lengthOfLISBinarySearch(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	d := make([]int, len(nums)) // 数组d的最大容量为len(nums)
	length := 0                 // 数组d的实际使用容量
	for _, x := range nums {    // 在长度length的数组d中，找到x的插入位置i
		i := binarySearchInsertPosition(d, length, x)
		d[i] = x         // 把数字x插入d[i]
		if i == length { // 如果i等于当前长度，说明x是追加到了数组d的尾部
			length++
		}
	}
	return length
}

// 二分搜索查找插入位置
func binarySearchInsertPosition(d []int, len, x int) int {
	low, high := 0, len-1
	for low <= high {
		mid := low + (high-low)/2
		if x < d[mid] {
			high = mid - 1
		} else if x > d[mid] {
			low = mid + 1
		} else {
			return mid
		}
	}
	return low
}
