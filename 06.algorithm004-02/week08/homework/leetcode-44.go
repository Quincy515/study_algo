package leetcode

// 动态规划 Time: O(m*n), Space: O(m*n)
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)   // s和p的长度,分别赋值给m和n,方便使用
	d := make([][]bool, m+1) // 定义辅助数组d
	for i := range d {
		d[i] = make([]bool, n+1)
	}

	// 1. 边界情况, s和p都是空字符串，两个空串是匹配的，所以 d(0,0)=true
	d[0][0] = true
	// 2. 处理模式串p为0的情况, d(i,0) = false 其中i>=1
	for i := 1; i <= m; i++ {
		d[i][0] = false
	}
	// 3. 处理s为空字符串的情况
	for j := 1; j <= n; j++ {
		if rune(p[j-1]) == '*' {
			d[0][j] = d[0][j-1]
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
				d[i][j] = d[i][j-1] || d[i-1][j]
			} else { // 对于所有其他情况d(i,j)为false
				d[i][j] = false
			}
		}
	}

	return d[m][n] // 填充完二维数组，返回d[m][n]即可
}

// 辅助函数进行判断sc和pc是否相等
func isEqual(sc, pc rune) bool {
	return sc == pc || pc == '?'
}

// 贪心算法 Time: O(m*n), Space: O(1)
func isMatchGreedy(s string, p string) bool {
	m, n := len(s), len(p) // s和p的长度,分别赋值给m和n,方便使用
	// 定义分别指向s和p的游标i和j,定义另外两个辅助变量sBegin, pBegin
	i, j, sBegin, pBegin := 0, 0, -1, -1
	for i < m { // 当游标i小于s的长度m时,重复以下操作
		if j < n && isEqual(rune(s[i]), rune(p[j])) {
			i, j = i+1, j+1 // 此时i和j游标的字符相等,游标向后移动
		} else if j < n && rune(p[j]) == '*' { // 表示一开始让'*'匹配0个字符
			j++                   // 否则如果p[j]为'*',先把j向后移动一个位置
			sBegin, pBegin = i, j // 把当前要开始对比的下标i和j分别记录到sBegin, pBegin
		} else if pBegin != -1 { // 没有进入前两个分支,查看p中是否有'*'
			sBegin++              // 让'*'去多匹配s中的一个字符
			i, j = sBegin, pBegin // sBegin加1赋值给i,游标j不用改变,即'*'的下一个位置
		} else { // 剩下的所有其他情况都直接返回false
			return false
		}
	}
	// 如果可以顺利退出循环,说明s字符串都已经匹配结束,需要考虑模式串p是否有剩余字符
	for j < n && rune(p[j]) == '*' { // 如果有,且p[j]为'*'就不断向后移动
		j++
	}
	return j == n // 最后判断j是否等于n即可,正好等于表明p要不在循环中正好匹配完,要么剩余一些字符但都是'*'
}

/**
S     P
aa    a     x
ac    *     ✔️
ac    *d    x
abcde *a*d? ✔️

简易正则表达式匹配
简易正则表达式匹配:  . 可以匹配任意单个非空字符
       通配符匹配:  ? 可以匹配任意单个非空字符
简易正则表达式匹配:  * 可以将它前面的一个字符重复0次或多次
       通配符匹配:  * 可以匹配0个字符或任意多个字符

动态规划思路
d(i,j) 表示 s(0...i-1) 和 p(0...j-1) 这两个前缀子串是否匹配
d(0,0) 边界情况, s和p都是空字符串，两个空串是匹配的，所以 d(0,0)=true
d(i,0) = false 其中i>=1
d(0,j) j>=1, 如果 p(j-1)=='*' 则 d(0,j)=d(0,j-1)
             如果 p(j-1)!='*' 则 d(0,j)=false

一般情况下的递推式
s 和 p 从 s(0...i-1) 和 p(0...j-1)是否匹配, 即d(i,j)情况
假设 sc := s[i-1] pc:=p[j-1]
如果 1. sc ⊜ pc 则 d(i,j) = d(i-1,j-1)
判断 sc 和 pc 是否相等，sc ⊜ pc { sc == pc || pc == '?' }
如果 2. pc == '*' 则 d(i,j) = { d(i,j-1) || d(i-1,j) }
       - '*' 匹配0个字符，相等于消除*号，即检测s(0...i-1)和p(0...j-2)是否匹配
       - '*' 匹配至少1个字符，即检测s(0...i-2)和p(0...j-1)是否匹配
如果 3. sc 和 pc 是两个不相等的小写字母，则两个前缀串一定不匹配 d(i,j) = false

d(m,n) m和n分别是字符串s和p的长度，
需要填充二维数组d，所以时间复杂度 Time: O(m*n),
使用辅助数组d，所以空间复杂度   Space: O(m*n),
可以使用两个一位数组进行滚动更新，也可以只使用一个一维数组加一个变量来保存状态，降低Space:O(n)

贪心算法思路
使用游标i和j来指向字符串s和p，i前面的s子串和j前面的p子串是匹配的
1. 如果 s[i] ⊜ p[j] 则i和j都向后移动一位i++,j++
2. 如果 p[j] == '*' 则p[j]这个'*'可以匹配s中0个或任意多个字符，
				并不知道匹配多少个才是正确的，贪心算法先匹配0个字符，失败再匹配1个字符，依次类推
游标i不变，游标j+1，记录当前s和p的位置，此时i和j左侧子串是匹配的，所以i和j当前的位置就是剩余要匹配的起始位置，记录sBegin=i, pBegin=j
3. 如果 pBegin != -1 表示已经遇到'*',则sBegin++,i=sBegin,j=pBegin
4. 否则 直接返回 false，表示s和p不匹配

j==n,i<m,pBegin!=-1  => 情况3或者情况4
i==n,j<n             => p[j] 后面都是'*'才匹配
Time: O(m*n) Space: O(1)
*/
