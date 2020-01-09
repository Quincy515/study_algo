package leetcode

// Time: O(n), Space: O(n)
func reverseWords(s string) string {
	if s == "" || len(s) == 0 { // 字符串为空或长度为0
		return s // 直接返回不需要处理
	}
	c := []rune(s)       // 把字符串转成字符数组
	start, end := 0, 0   // 定义开始和结束游标并初始化为0
	for start < len(c) { // 当开始游标小于字符长度时,执行以下操作
		// 当结束游标小于字符长度并且结束游标指向的字符不是空格
		for end < len(c) && c[end] != ' ' {
			end++ // 向后移动一位
		}
		// 跳出循环翻转start到end-1的子串
		for i, j := start, end-1; i < j; i, j = i+1, j-1 {
			c[i], c[j] = c[j], c[i]
		}
		// 然后把开始游标和结束游标到移动到end的下一个位置
		start = end + 1
		end = start
	}
	// 循环结束后用处理过的字符数组构建新的字符串返回即可
	return string(c)
}
