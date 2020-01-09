package leetcode

// Time: O(n) Space: O(1)
func isPalindrome(s string) bool {
	if s == "" || len(s) == 0 {
		return true // 边界情况，字符串为空或者长度为0，确定是回文字符串
	}
	i, j := 0, len(s)-1
	for ; i < j; i, j = i+1, j-1 {
		for i < j && !isAlphanumeric(rune(s[i])) {
			i++
		}
		for i < j && !isAlphanumeric(rune(s[j])) {
			j--
		}
		if i < j && !isEqualIgnoreCase(rune(s[i]), rune(s[j])) {
			return false
		}
	}
	return true
}

// 辅助函数 判断字符是否是字母或数字
func isAlphanumeric(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

// 辅助函数 判断忽略大小写后是否相等
func isEqualIgnoreCase(a, b rune) bool {
	if a >= 'A' && a <= 'Z' {
		a += 32
	}
	if b >= 'A' && b <= 'Z' {
		b += 32
	}
	return a == b
}
