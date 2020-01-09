package leetcode

// 动态规划 Time: O(m*n), Space: O(m*n)
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)   // s和p的长度,分别赋值给m和n,方便使用
	d := make([][]bool, m+1) // 辅助数组d
	for i := 0; i <= m; i++ {
		d[i] = make([]bool, n+1)
	}
	// 1. 边界情况, s和p都是空字符串，两个空串是匹配的，所以 d(0,0)=true
	d[0][0] = true
	// 2. 处理模式串p为空的情况, d(i,0) = false 其中i>=1
	for i := 1; i <= m; i++ {
		d[i][0] = false
	}
	// 3. 处理s为空字符串的情况
	for j := 1; j <= n; j++ {
		if rune(p[j-1]) == '*' {
			d[0][j] = d[0][j-2]
		} else {
			d[0][j] = false
		}
	}
	// 4. 处理一般情况下的d(i,j)
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// 当前考察的两个前缀串最后一个字符取出
			sc, pc := rune(s[i-1]), rune(p[j-1])
			if isEqual(sc, pc) { // 如果相等
				d[i][j] = d[i-1][j-1]
			} else if pc == '*' { // 否则pc为'*'
				// 取出pc前一个字符定义为preChar
				preChar := rune(p[j-2])
				if isEqual(sc, preChar) {
					d[i][j] = d[i][j-2] || d[i][j-1] || d[i-1][j]
				} else {
					d[i][j] = d[i][j-2]
				}
			} else { // 对于所有其他情况d(i,j)为false
				d[i][j] = false
			}
		}
	}

	return d[m][n] // 填充完二维数组，返回d[m][n]即可
}

// 辅助函数 判断sc和pc是否相等⊜
func isEqual(sc, pc rune) bool {
	return sc == pc || pc == '.'
}

// 两个滚动一维数组方法 动态规划 Time: O(m*n), Space: O(n)
func isMatchTwoArray(s string, p string) bool {
	m, n := len(s), len(p) // s和p的长度,分别赋值给m和n,方便使用
	pre, cur := make([]bool, n+1), make([]bool, n+1)
	// 1. 边界情况, s和p都是空字符串，两个空串是匹配的，所以 d(0,0)=true
	pre[0] = true
	// 2. 处理模式串p为空的情况, d(i,0) = false 其中i>=1
	for i := 1; i <= m; i++ {
		pre[0] = false
	}
	// 3. 处理s为空字符串的情况
	for j := 1; j <= n; j++ {
		if rune(p[j-1]) == '*' {
			pre[j] = pre[j-2]
		} else {
			pre[j] = false
		}
	}
	// 4. 处理一般情况下的d(i,j)
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// 当前考察的两个前缀串最后一个字符取出
			sc, pc := rune(s[i-1]), rune(p[j-1])
			if isEqual(sc, pc) { // 如果相等
				cur[j] = pre[j-1]
			} else if pc == '*' { // 否则pc为'*'
				// 取出pc前一个字符定义为preChar
				preChar := rune(p[j-2])
				if isEqual(sc, preChar) {
					cur[j] = cur[j-2] || cur[j-1] || pre[j]
				} else {
					cur[j] = cur[j-2]
				}
			} else { // 对于所有其他情况d(i,j)为false
				cur[j] = false
			}
		}
		// 求出最新状态存储在cur数组中
		// cur和pre数组交换
		// 交换后pre就指向最新的状态数组
		// cur指向旧数组,在下一轮计算中得到的新值,存储在cur中
		cur, pre = pre, cur
		for i := range cur {
			cur[i] = false // cur指向的旧数组值全部修改为false
		} // 在下一轮迭代中使用
	}

	return pre[n] // pre指向最新的状态数组，返回pre[n]即可
}
