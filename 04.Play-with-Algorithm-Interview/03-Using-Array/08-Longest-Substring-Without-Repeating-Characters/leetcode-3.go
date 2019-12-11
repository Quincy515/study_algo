package LeetCode

// 滑动窗口 时间复杂度: O(len(s)) 空间复杂度: O(len(charset))
func lengthOfLongestSubstring1(s string) int {
	freq := make([]int, 256)
	l, r := 0, -1 // 滑动窗口为s[l...r]
	res := 0      //最长子串的长度

	// 整个循环从l==0;r==-1这个空窗口开始
	// 到l==len(s);r==len(s)-1这个空窗口截止
	// 在每次循环里逐渐改变窗口,维护freq,并记录当前窗口中是否找到一个新的最优值
	for l < len(s) { // 窗口的左边界在数组范围内,则循环继续
		if r+1 < len(s) && freq[s[r+1]] == 0 {
			r++
			freq[s[r]]++
		} else { // r已经到头|| freq[s[r+1]] == 1
			freq[s[l]]--
			l++
		}
		res = max(res, r-l+1)
	}
	return res
}

// 滑动窗口代码优化 时间复杂度: O(len(s)) 空间复杂度: O(len(charset))
func lengthOfLongestSubstring2(s string) int {
	freq := make([]int, 256)
	l, r := 0, -1 // 滑动窗口为s[l...r]
	res := 0

	// 在这里,循环终止条件可以是r+1<len(s)
	for r+1 < len(s) {
		if freq[s[r+1]] == 0 {
			r++
			freq[s[r]]++
		} else { // freq[r+1] == 1
			freq[s[l]]--
			l++
		}
		res = max(res, r-l+1)
	}
	return res
}

// 滑动窗口另一种实现,仅做参考 时间复杂度: O(len(s)) 空间复杂度: O(len(charset))
func lengthOfLongestSubstring3(s string) int {
	freq := make([]int, 256)
	l, r := 0, -1 // 滑动窗口为s[l...r]
	res := 0

	for r+1 < len(s) {
		for r+1 < len(s) && freq[s[r+1]] == 0 {
			r++
			freq[s[r]]++
		}
		res = max(res, r-l+1)

		if r+1 < len(s) {
			r++
			freq[s[r]]++
			for l <= r && freq[s[r]] == 2 {
				freq[s[l]]--
				l++
			}
		}
	}
	return res
}

// l每次可以向前跳跃,而不仅仅是+1,但代价是,为了获得这个跳跃位置,每次需要遍历整个窗口的字符串
// 时间复杂度: O(len(s)*len(charset)) 空间复杂度: O(1)
func lengthOfLongestSubstring4(s string) int {
	l, r := 0, 0 // 滑动窗口为s[l...r]
	res := 0
	for r < len(s) {
		index := isDuplicateChar(s, l, r)
		// 如果s[r]之前出现过
		// l可以直接跳到s[r+1]之前出现的位置+1的地方
		if index != -1 {
			l = index + 1
		}
		res = max(res, r-l+1)
		r++
	}
	return res
}

// 查看s[l...r-1]之间是否存在s[r],若存在,返回相应的索引,否则返回-1
func isDuplicateChar(s string, l, r int) int {
	for i := l; i < r; i++ {
		if s[i] == s[r] {
			return i
		}
	}
	return -1
}

// 滑动窗口 其中使用last[c]保存字符c上一次出现的位置,用于在右边界发现重复字符时,快速移动左边界
// 使用这种方法,时间复杂度依然为O(n),但是只需要动r指针,实际上对整个s只遍历一次
// 相较而言,之前的方法,需要移动l和r两个指针,相对于对s遍历了两次
// 时间复杂度: O(len(s)) 空间复杂度: O(len(charset))
func lengthOfLongestSubstring5(s string) int {
	last := make([]int, 256)
	for i := range last {
		last[i] = -1
	}
	l, r := 0, -1 // 滑动窗口为s[l...r]
	res := 0
	for r+1 < len(s) {
		r++
		if last[s[r]] != -1 {
			l = max(l, last[s[r]]+1)
		}
		res = max(res, r-l+1)
		last[s[r]] = r
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstring6(s string) int {
	lastOccurred := make(map[rune]int) // 记录已经访问的字母
	l := 0                             // 左边界的指针
	res := 0                           // 返回最长子串的长度
	for i, c := range []rune(s) {      // 遍历字符串
		if lastIndex, ok := lastOccurred[c]; ok && lastIndex >= l {
			l = lastIndex + 1
		}
		if i-l+1 > res {
			res = i - l + 1
		}
		lastOccurred[c] = i
	}
	return res
}
