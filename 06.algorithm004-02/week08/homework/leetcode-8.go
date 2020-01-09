package leetcode

import "math"

func myAtoi(str string) int {
	start := 0        // 有效数字的初始下标初始化为0
	p := 0            // 表示遍历字符串的游标
	n := len(str)     // 表示字符串的长度
	negative := false // 表示字符串是否为负数
	for p < n && rune(str[p]) == ' ' {
		// 忽略字符的前导空格
		p++
	}
	if p == n { // 全部为空格或空串
		return 0 // 不能转为整数返回0
	}

	if rune(str[p]) == '+' {
		p++ // 游标所指向的字符为加号,继续移动
	} else if rune(str[p]) == '-' {
		p++
		negative = true // 记录数字的正负性
	}

	for p < n && rune(str[p]) == '0' {
		p++ // 考虑数字部分可能开始为0,去除前导0
	}
	start = p // 记录不为0的开始下标

	for p < n && rune(str[p]) >= '0' && rune(str[p]) <= '9' {
		p++ // 游标p扫过连续的数字
	}
	if p == start { // 表示数字部分一个数字都不包含
		return 0 // 不能合法的转为整数
	}

	if p-start > 10 {
		if negative {
			return math.MinInt32
		} else {
			return math.MaxInt32
		}
	}

	var num int // 把字符串的数字部分先转为num
	for i := start; i < p; i++ {
		num = num*10 + int(rune(str[i])-'0')
	}
	if negative { // 记录的数字是负数
		num = -num // 加上负号
	}
	// 转成的数字小于int的最小值或者大于int的最大值
	if num < math.MinInt32 {
		return math.MinInt32 // 返回相应的极值
	} else if num > math.MaxInt32 {
		return math.MaxInt32
	}
	return int(num)
}
