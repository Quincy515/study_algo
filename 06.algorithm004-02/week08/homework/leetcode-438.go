package leetcode

// 只使用一个辅助数组 Time:O(n) Space: O(k),n为字符串s的长度，k为字符集的大小
func findAnagramsOn(s string, p string) []int {
	// 先定义保存开始下标的结果列表
	var result []int
	// 判断边界条件,s或p为空,或s的长度小于p的长度
	if s == "" || p == "" || len(s) < len(p) {
		return result // 直接返回空结果
	}
	// 否则，先把s和p的长度取出来赋值给变量方便使用
	sLen, pLen := len(s), len(p)
	// 定义大小为26的辅助数组pc
	pc := make([]rune, 26)

	// pc中保存了字符串p的字符数量统计
	for i := 0; i < pLen; i++ {
		// 把P中下标为I的字符取出减去字符'a'
		// 然后把pc中对应位置的数量加一
		pc[rune(p[i])-'a']++
	}
	// 定义滑动窗口的左右边界left和right初始都为0
	left, right := 0, 0
	for right < sLen { // 当右边界小于s的长度的时
		// 先检查右边界指向的字符在数组pc中数量是否大于0
		if pc[rune(s[right])-'a'] > 0 {
			// 是的话就把pc中这个字符的数量减一
			// 相当于用滑动窗口中的字符抵消它
			pc[rune(s[right])-'a']--
			// 然后右边界加一向右移动
			right++
		} else { // 否则说明右边界指向的字符在pc中的数量已经不大于0
			// 要么这个字符已经被抵消完
			// 要么出现一个不在p中的字符
			// 这时我们把左边界指向的字符在pc中的数量加回来
			pc[rune(s[left])-'a']++
			// 然后左边界向右移动一个位置
			left++
		}
		// 右边界或左边界发生移动后,就检查一下滑动窗口的长度
		if right-left == pLen {
			// 是的话，说明找到一个异位词，把开始下标记录到结果列表
			result = append(result, left)
		}
	}
	return result // 循环结束后返回结果列表即可
}

// 辅助函数 用于判断辅助数组SC和PC中包含的各个字符对应的数量是否相同
func equals(sc, pc []rune) bool {
	// 下标i从0遍历到sc的最后一个下标
	for i := 0; i < len(sc); i++ {
		if sc[i] != pc[i] {
			return false
		}
	}
	// 循环结束后没有提前返回，说明sc和pc中相同字符对应的数量是一样的
	return true
}

// 使用两个辅助数组 Time: O(n*k) Space: O(k),n为字符串s的长度，k为字符集的大小
func findAnagrams(s string, p string) []int {
	// 先定义保存开始下标的结果列表
	var result []int
	// 判断边界条件,s或p为空,或s的长度小于p的长度
	if s == "" || p == "" || len(s) < len(p) {
		return result // 直接返回空结果
	}
	// 否则，先把s和p的长度取出来赋值给变量方便使用
	sLen, pLen := len(s), len(p)
	// 定义两个大小为26的辅助数组pc和sc
	// 这里对于每个小写字母，使用一个rune来表示它的计数
	pc, sc := make([]rune, 26), make([]rune, 26)
	// 由于题目说明s和p的长度都不会超过100，所以一个rune足够存储了，不必担心溢出

	for i := 0; i < pLen; i++ {
		// 把P中下标为I的字符取出减去字符'a'
		// 然后把pc中对应位置的数量加一
		pc[rune(p[i])-'a']++
		// 同样的把s中下标为i的字符取出减去字符'a'
		// 然后把sc中对应位置的数量加一
		sc[rune(s[i])-'a']++
	}
	// 循环结束后，pc中就保存了字符串p的字符数量统计
	// 而sc中保存了第一个滑动窗口内的字符数量统计
	// 判断一下它们是否相等
	// 如果相等就说明找到一个p的异位词
	// 把开始下标0加入结果列表
	if equals(sc, pc) {
		result = append(result, 0)
	}
	// 接下来开始i从pLen开始一直遍历到sLen-1
	for i := pLen; i < sLen; i++ {
		// 右边加入滑动窗口的字符是s[i]
		// 于是s[i]减去字符a后对应的位置计数加一
		sc[rune(s[i]-'a')]++
		// 左边离开滑动窗口的字符对应的下标是i-pLen
		// 这个字符减去'a'后对应的位置计数减一
		sc[rune(s[i-pLen])-'a']--
		// 接着只要判断此时的sc是否等于pc
		// 是的话说明找到p的一个异位词
		// 当前s的子串是s[i-pLen+1,i],长度是pLen
		// 把开始下标i-pLen+1加入到结果列表
		if equals(sc, pc) {
			result = append(result, i-pLen+1)
		}
	}
	// 循环结束后返回结果列表即可
	return result
}
