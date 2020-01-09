package leetcode

func reverseWords151(s string) string {
	if s == "" || len(s) == 0 {
		return s // 字符串为空,或长度为0就返回它自己
	}
	str := []rune(s) // 把字符串转为字符数组
	// 定义游标p和q，并初始化end指向字符串末尾
	p, q, end := 0, 0, len(str)-1
	for end >= 0 && str[end] == ' ' {
		end-- // 当end为空格时不断向前移动
	}
	for q <= end {
		start := p // 记录单词开始下标
		for q <= end && str[q] == ' ' {
			q++ // 游标q跳过空格
		}
		for q <= end && str[q] != ' ' {
			str[p] = str[q] // q不指向空格,就把字符拷贝到游标p的位置
			p, q = p+1, q+1 // 同时把p和q都向后移动一位
		}
		// 拷贝完一个单词后去翻转这个单词
		reverse(str, start, p-1)
		if q <= end { // 当q还没有移动到end时,说明还有单词要处理
			str[p] = ' ' // 在p的位置插入一个空格
			p++          // 并向后移动一位
		}
	}
	// 循环结束后，把字符串中有效部分翻转
	reverse(str, 0, p-1)
	return string(str[0:p])
}

// 辅助函数，用于翻转字符数组中i到j之间的字符
func reverse(str []rune, i, j int) {
	for ; i < j; i, j = i+1, j-1 {
		// 当i小于j时就交换i,j指向的字符
		str[i], str[j] = str[j], str[i]
		// 游标i和j同时向中间移动一位
	}
}
