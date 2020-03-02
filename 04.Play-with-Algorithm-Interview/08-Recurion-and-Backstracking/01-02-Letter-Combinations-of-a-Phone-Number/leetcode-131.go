package leetcode

// 辅助分隔函数，输入字符串s，每次递归分隔字符串的开始下标,表示任意子串是否回文的状态数组d,最后结果数组result,每次分隔的子串集合elem
func partitionHelp(s string, start int, d [][]bool, result *[][]string, elem []string) {
	if start >= len(s) { // 如果开始下标已经大于等于字符串长度
		tmp := make([]string, len(elem))
		copy(tmp, elem)                // 说明产生一个有效的分隔
		*result = append(*result, tmp) // 把elem加到结果里
	} else { // 否则子串的结束下标end从start开始一直遍历到字符串最后
		for end := start; end < len(s); end++ {
			if d[start][end] { // 如果下标start到end之间的子串是一个回文串
				elem = append(elem, s[start:end+1])      // 则把子串取出来并加入到当前子串的分隔集合
				partitionHelp(s, end+1, d, result, elem) // 然后递归的调用自己,开始下标从end+1开始
				elem = elem[:len(elem)-1]                // 递归返回后则需要把最后加入的回文子串移除
				// 下标end会在循环中+1，尝试下一个分隔
			}
		}
	}
}

// Time: O(2^n), Space: O(n^2)
func partition(s string) [][]string {
	var result [][]string       // 定义要返回的结果
	if s == "" || len(s) == 0 { // 当字符串为空或长度为0时
		return result //直接返回空结果
	}
	n := len(s)            // 否则把字符串的长度赋值给n
	d := make([][]bool, n) // 定义状态数组d用于表示任意子串是否是回文串
	for i := range d {
		d[i] = make([]bool, n)
	}
	// 接下来用两层循环来填充这个状态数组,状态数组的详解见647.回文子串
	for i := n - 1; i >= 0; i-- { // i从大到小遍历
		for j := i; j < n; j++ { // j从i开始，从小到大遍历
			if i == j { // 当前子串只有一个字符
				d[i][j] = true // 因此是回文串
			} else if i+1 == j { // i和j是相邻的两个字符
				d[i][j] = false   // 默认这两个字符不是回文串
				if s[i] == s[j] { // 判断这两个字符是否相等
					d[i][j] = true // 如果相等则说明两个字符是回文串
				}
			} else { // 其他情况不仅需要i和j指向的字符相等
				// 而且除去这两个字符的中间子串也要是回文串
				d[i][j] = false // 默认该子串不是回文串
				if s[i] == s[j] && d[i+1][j-1] {
					d[i][j] = true // 满足这两个条件说明改子串是回文串
				}
			}
		}
	}
	// 准备工作做完后，直接调用辅助的递归分隔函数即可
	partitionHelp(s, 0, d, &result, []string{})
	return result // 最后返回结果
}
