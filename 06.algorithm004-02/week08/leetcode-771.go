package leetcode

// 暴力破解版本 Time:O(m*n), Space: O(1)
func numJewelsInStones(J string, S string) int {
	if len(S) == 0 || len(J) == 0 {
		return 0
	}
	cnt := 0                      // 初始化珠宝数量为0
	for i := 0; i < len(J); i++ { // 遍历珠宝种类
		for j := 0; j < len(S); j++ { // 对于一种具体的珠宝,遍历字符串s
			if J[i] == S[j] { // 如果S中当前字符是珠宝
				cnt++ // 珠宝数量加一
			}
		}
	}
	return cnt // 返回珠宝数量即可
}

// 使用哈希表 Time: O(m+n), Space: O(k) k为字符集大小
func numJewelsInStonesHash(J string, S string) int {
	if len(S) == 0 || len(J) == 0 {
		return 0 // 处理边界情况
	}
	// 这里使用ASCII字符集大小256
	var d [256]bool               // 使用bool数组模拟哈希表
	cnt := 0                      // 初始化珠宝数量为0
	for i := 0; i < len(J); i++ { // 遍历珠宝种类
		d[J[i]] = true
	}
	for j := 0; j < len(S); j++ { // 对于一种具体的珠宝,遍历字符串s
		if d[S[j]] {
			cnt++
		}
	}
	return cnt
}

func numJewelsInStonesMap(J string, S string) int {
	count := 0
	mark := map[rune]bool{}
	for _, c := range J {
		mark[c] = true
	}
	for _, s := range S {
		if mark[s] {
			count++
		}
	}
	return count
}
