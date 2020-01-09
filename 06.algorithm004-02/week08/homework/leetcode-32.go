package leetcode

// 使用辅助栈 Time:O(n), Space: O(n), Time Limit Exceeded
func longestValidParentheses(s string) int {
	if s == "" || len(s) == 0 {
		return 0 // 处理边界情况，字符串为空，或长度为0
	}
	n := len(s)            // 把字符串长度赋值给n
	p := -1                // 游标p模拟栈顶指针
	max := 0               // 记录最大长度
	st := make([]int, n+1) // 定义长度为n+1的数组来模拟栈
	p++                    // 同时栈顶指针p+1变为0
	st[p] = -1             // -1入栈

	for i := 0; i < n; i++ { // 遍历字符串
		// 栈顶指针不为-1, 栈顶指针对应的字符是(, 且当前字符是)
		if st[p] != -1 && []rune(s)[st[p]] == '(' && []rune(s)[i] == ')' {
			p--                // p-1表示将栈顶元素出栈
			if i-st[p] > max { // 计算当前子串长度,i-当前栈顶指针表示的下标
				max = i - st[p] // 取较大值更新max
			}
		} else { // 当前下标入栈
			p++ // 栈顶指针+1
			st[p] = i
		}
	}
	return max // 循环结束后返回max即可
}

// 定义状态d(i)表示以第i个字符结尾的有效括号子串的最大长度
// 动态规划 Time:O(n), Space: O(n)
func longestValidParenthesesDP(s string) int {
	if s == "" || len(s) == 0 {
		return 0 // 处理边界情况，字符串为空，或长度为0
	}
	n := len(s)         // 把字符串长度赋值给n
	left := 0           // 定义左边括号数量
	max := 0            // 记录最大长度
	d := make([]int, n) // 定义大小为n的状态数组

	for i := 0; i < n; i++ { // 遍历字符串
		if []rune(s)[i] == '(' { // 当前字符为左括号
			left++ // 左括号数量+1
		} else if left > 0 { // 否则当前字符为右括号且左括号数量大于0,表示能形成有效括号对
			d[i] = d[i-1] + 2 // d[i]初始化为前一状态d[i-1]+新增一对括号的长度2
			if i-d[i] >= 0 {  // 若当前子串前一位置i-d[i]是否>=0
				d[i] += d[i-d[i]] // 加上对应的状态值
			} // 用于处理相邻的有效括号子串
			if d[i] > max { // max和d[i]中的较大值更新max
				max = d[i]
			}
			left-- // 左括号数量-1
		}
	}
	return max // 循环结束后返回max即可
}
