package leetcode

// 动态规划 Time: O(O^2) Space: O(n)
func longestPalindrome(s string) string {
	if s == "" || len(s) == 0 {
		// 当字符串为空或长度为0时返回空串
		return ""
	}
	// 初始化字符串长度为n,最长回文子串的开始下标为start,长度为maxLen
	n, start, maxLen := len(s), 0, 0
	// 定义二维数组d保存s(i,j)是否为回文字符串
	d := make([][]bool, n)
	for i := 0; i < n; i++ {
		d[i] = make([]bool, n)
	}

	// 使用两个for循环来求s(i,j)是否为回文字符串
	for i := n - 1; i >= 0; i-- { // i从大到小遍历
		for j := i; j < n; j++ { // j从小到大遍历
			if i == j { // 只有一个字符
				d[i][j] = true
			} else if i+1 == j { // 有两个字符，判断i和j指向的字符是否相等
				d[i][j] = rune(s[i]) == rune(s[j])
			} else { // 其他情况不仅需要i和j指向的字符相等
				// 而且除开这两个字符的中间子串也要是回文字符串才能得出当前子串为回文串
				d[i][j] = rune(s[i]) == rune(s[j]) && d[i+1][j-1]
			}

			// 如果当前子串s(i,j)为回文串，且j-i+1大于最大长度
			if d[i][j] && (j-i+1) > maxLen {
				// 则更新子串的开始下标start和最大长度
				start = i
				maxLen = j - i + 1
			}
		}
	}

	// 循环结束后只需要，把开始下标start长度为maxLen的子串返回即可
	return s[start : start+maxLen]
}

// 从中心向两边扩展回文子串的方法 Time:O(n^2) Space:O(1)
func longestPalindromeExpand(s string) string {
	if s == "" || len(s) == 0 {
		// 当字符串为空或长度为0时返回空串
		return ""
	}
	// 初始化最长回文子串的开始下标为start,长度为maxLen
	start, maxLen := 0, 0

	for i := 0; i < len(s); i++ { // 遍历字符串
		// 以第i个字符为中心扩展子串，得到的最长回文子串长度为len1
		len1 := expand(s, i, i)
		// 以第i个和第i+1这两个字符为中心扩展子串，得到的最长回文子串长度为len2
		len2 := expand(s, i, i+1)
		len := len1
		if len1 < len2 {
			len = len2
		}
		if len > maxLen {
			start = i - (len-1)/2
			maxLen = len
		}
	}
	return s[start : start+maxLen]
}

// 扩展函数,输入字符串和向两边扩展的左右游标,输出是可以扩展的最长回文子串长度
func expand(s string, left, right int) int {
	// 当左右游标不超过左右边界，且指向的字符相等时，将游标向两边移动
	for left >= 0 && right < len(s) && rune(s[left]) == rune(s[right]) {
		left, right = left-1, right+1
	}
	// 最后左右游标不满足边界限制或者指向的字符不相等时，循环结束
	// 此时回文子串的右边界为right-1,左边界为left+1
	// 此时回文字符串的长度为 right-1 - (left+1)-1
	return right - left - 1
}
