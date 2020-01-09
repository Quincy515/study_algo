package leetcode

// Time: O(k*n) Space: O(1) k为n个字符串最长公共前缀的长度
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return "" // 字符串长度为0,返回空字符串
	}
	first := strs[0]                  // 把第一个字符串取出作为基准字符串
	for i := 0; i < len(first); i++ { // 用下标i遍历基准字符串
		c := rune(first[i])              // 先取出基准字符串在下标i上的字符c
		for j := 1; j < len(strs); j++ { // 然后用j遍历数组中剩余的字符串
			// 如果下标i超出了当前字符串的长度,或者当前字符串在下标i上的字符不等于c
			if i >= len(strs[j]) || rune(strs[j][i]) != c {
				return first[:i] // 结束操作,返回基准字符串中0到i的子串即可
			}
		}
	}
	return first // 如果循环结束后,还未提前返回,说明遍历完第一个字符串,则基准字符串即为最长公共前缀
}
