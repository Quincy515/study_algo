package leetcode

// 动态规划 Time: O(n^2), Space: O(n^2)
func countSubstrings(s string) int {
	// 边界情况,当字符串为空或长度为0时
	if s == "" || len(s) == 0 {
		return 0 // 直接返回0
	}
	// 初始字符串长度为n,回文子串的个数为0
	n, count := len(s), 0
	d := make([][]bool, n) // 定义二维数组d来保存子串s(i,j)是否为回文字符串
	for i := range d {
		d[i] = make([]bool, n)
	}

	// 使用两层循环求s(i,j)是否为回文字符串
	for i := n - 1; i >= 0; i-- { // i从大到小遍历
		for j := i; j < n; j++ { // j从i开始,从小到大遍历
			if i == j { // 当前子串只有一个字符
				d[i][j] = true // 因此是回文串
			} else if i+1 == j { // i和j相邻只要两个字符
				d[i][j] = false
				if s[i] == s[j] { // 判断这两个字符是否相等
					d[i][j] = true // 即可判断当前子串是否为回文串
				}
			} else { // 其他情况不仅需要i和j指向的字符相等
				// 而且除开这两个字符的中间子串也要是回文字符串
				d[i][j] = false
				if s[i] == s[j] && d[i+1][j-1] {
					d[i][j] = true
				}
			}

			if d[i][j] { // 如果s(i,j)是回文字符串
				count++ // 则count+1
			}
		}
	}
	return count // 循环结束后返回count即可
}

// 辅助函数 定义扩展函数，
// 输入是一个字符串，和两个向两边扩展的游标
// 输出是扩展过程中，回文子串的个数
func expand(s string, left, right int) int {
	count := 0 // 定义回文子串的个数为0
	// 当左右游标不超过左右边界并且它们指向的字符相等时
	for left >= 0 && right <= len(s)-1 && s[left] == s[right] {
		count++                       // 说明两个游标中间的子串是回文子串，count+1
		left, right = left-1, right+1 // 两个游标向两边移动
	}
	return count // 循环结束后返回count即可
}

// 从中心向两边扩展回文子串的方法 Time: O(n^2), Space: O(1)
func countSubstrings(s string) int {
	// 边界情况,当字符串为空或长度为0时
	if s == "" || len(s) == 0 {
		return 0 // 直接返回0
	}
	// 初始回文子串的个数为0
	count := 0
	for i := 0; i < len(s); i++ { // 遍历字符串
		// 以第i个字符为中心扩展子串
		// 得到的回文子串数量加到count上
		count += expand(s, i, i)
		// 类似的以第i和i+1这两个字符为中心扩展子串
		// 得到的回文子串数量也加到count上
		count += expand(s, i, i+1)
	}
	return count // 循环结束后返回count即可
}
