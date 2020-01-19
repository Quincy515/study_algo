package leetcode

// Time: O(n), Space: O(m)
func minWindow(s string, t string) string {
	// 处理边界情况，如果s或t为空
	if s == "" || t == "" {
		return "" // 则返回空字符串
	}
	// 定义一个数组来记录需要寻找到的字符对应的数量
	// 只考虑ASCII字符，定义大小256的数组来模拟哈希表
	required := [256]int{}
	// 定义要寻找到的最短子串的开始下标和长度
	// 以及需要找到的字符总量即t的长度
	start, length, requiredCnt := 0, len(s)+1, len(t)
	// 定义滑动窗口的左右边界,初始化为下标0
	left, right := 0, 0
	// 接着遍历一遍字符串t
	for i := 0; i < len(t); i++ {
		// 把需要寻找的字符以及相应的数量存储在数组中
		required[t[i]]++
	}
	// 窗口右边界从0遍历到字符串s的右边界
	for ; right < len(s); right++ {
		r := s[right]        // 每到一个位置就取出窗口右边界的字符
		if required[r] > 0 { // 如果在数组中的数量大于0
			// 说明找到了一个t中的字符
			requiredCnt-- // 需要寻找到的字符总量-1
		}
		required[r]-- // 字符数组中对应的字符数量-1

		for requiredCnt == 0 { // 当需要寻找到的字符总量等于0
			// 说明当前滑动窗口内包含了t中所有的字符
			// 执行以下操作，如果当前滑动窗口的长度小于length
			if right-left+1 < length {
				// 就更新最短子串的开始下标和长度
				start = left
				length = right - left + 1
			}
			// 接着移动滑动窗口左边界
			l := s[left]  // 先把左边界的字符取出
			required[l]++ // 把它在数组中的数量+1
			// 表示它已经被排除到滑动窗口外
			if required[l] > 0 { // 如果此时这个字符的数量大于0
				// 说明还需要往窗口中添加字符才能使该窗口再次包含t中所有字符
				requiredCnt++ // 因此所需要的字符总量+1
			}
			left++ //窗口左边界+1表示向右移动
		}
	}
	// 循环结束长度依然为字符串s的长度+1
	if length == len(s)+1 {
		return "" // 说明没有找到满足条件的子串,返回空串即可
	}
	return s[start : start+length] // 否则根据开始下标和长度返回找到的最短子串
}
