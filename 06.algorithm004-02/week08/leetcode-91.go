package leetcode

// 辅助函数 判断数字字符a，b能否有效解码
func isValidTwoDigit(a, b rune) bool {
	return (a == '1' && b <= '9') || (a == '2' && b <= '6')
}

// 辅助递归函数 输入字符数组c和开始下标i 输出是有效的解码方式数量
func decode(c []rune, i int) int {
	// 先处理递归终止条件
	if i == len(c) { // i正好等于数组长度
		return 1 // 找到有效解码返回1
	}
	// i正好指向数组最后一个字符并且不等于指向的字符不等于'0'
	if i == len(c)-1 && c[i] != '0' {
		return 1 // 说明找到有效解码返回1
	}
	if i > len(c) { // i大于数组长度
		return 0 // 不是有效解码返回0
	}
	num := 0         // 初始化解码方式数量为0
	if c[i] != '0' { // 如果当前单个字符可以有效解码
		num += decode(c, i+1) // 加上i+1开始的解码方式数量
	}
	// 如果i和i+1字符组合可以有效解码
	if i+1 < len(c) && isValidTwoDigit(c[i], c[i+1]) {
		num += decode(c, i+2) // 加上i+2开始的解码方式数量
	}
	return num // 最后返回解码方式数量
}

// 递归方法
func numDecodingsRecursive(s string) int {
	return decode([]rune(s), 0)
}

// Time: O(n), Space: O(n) n为字符串长度
func numDecodingsDP(s string) int {
	// 因为题目定义了非空有效的数字字符串,所以不用判断边界情况
	d := make([]int, len(s)+1) // 定义状态数组d
	d[0] = 1                   // 初始d(0)等于1
	d[1] = 0                   // d(1)根据s[0]是否能有效解码
	if rune(s[0]) != '0' {
		d[1] = 1
	}

	// i从2遍历到字符串长度
	for i := 2; i <= len(s); i++ {
		if rune(s[i-1]) != '0' { // 当前字符s[i-1]能有效解码
			d[i] += d[i-1]
		} // 如果i-2到i-1的字符组合能有效解码
		if isValidTwoDigit(rune(s[i-2]), rune(s[i-1])) {
			d[i] += d[i-2]
		}
	}

	// 循环结束后返回状态数组最后的元素即可
	return d[len(s)]
}

// Time: O(n), Space: O(1) n为字符串长度
func numDecodingsDPO1(s string) int {
	first := 1  // 初始d(0)等于1
	second := 0 // d(1)根据s[0]是否能有效解码
	if rune(s[0]) != '0' {
		second = 1
	}

	// i从2遍历到字符串长度
	for i := 2; i <= len(s); i++ {
		third := 0               // 循环中定义当前状态为third初始化为0
		if rune(s[i-1]) != '0' { // 当前字符s[i-1]能有效解码
			third += second
		} // 如果i-2到i-1的字符组合能有效解码
		if isValidTwoDigit(rune(s[i-2]), rune(s[i-1])) {
			third += first
		}
		// 滚动更新
		first = second
		second = third
	}

	// 循环结束后返回状态数组最后的元素即可
	return second
}
