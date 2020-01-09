package leetcode

// Time: O(n) Space: O(m) m为字符集大小
func firstUniqChar1(s string) int {
	var freq [26]int

	for _, v := range s {
		freq[v-'a']++
	}

	for k, v := range s {
		if freq[v-'a'] == 1 {
			return k
		}
	}
	return -1
}

// 遍历一遍数组 Time: O(n) Space: O(m)
func firstUniqChar(s string) int {
	if s == "" || len(s) == 0 {
		return -1 // 处理边界情况
	}
	var count [26]int
	var pos [26]int // 定义存储下标的pos数组
	for i := range pos {
		pos[i] = -1 // pos数组存储的下标初始化为-1
	}

	for i, v := range s { // 遍历字符串
		idx := v - 'a' // 当前字符减'a'得到下标位置
		count[idx]++   // count中对应下标数量加一
		if pos[idx] == -1 {
			// 当前字符第一次出现
			pos[idx] = i // 记录下标
		}
	}
	firstPos := len(s) + 1 // 记录只出现一次的字符中下标最小的
	for i := 0; i < 26; i++ {
		if count[i] == 1 {
			if firstPos > pos[i] {
				firstPos = pos[i]
			}
		}
	}
	if firstPos == len(s)+1 {
		return -1
	}
	return firstPos
}
