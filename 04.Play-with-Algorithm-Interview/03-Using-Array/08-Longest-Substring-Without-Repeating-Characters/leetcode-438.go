package LeetCode

// Sliding Window Time Complexity: O(len(p) + len(s)) Space Complexity: O(1)
func findAnagrams1(s string, p string) []int {
	l, r := 0, -1 // 滑动窗口为s[l...r]
	freq_s := make([]int, 26)
	freq_p := make([]int, 26)
	res := make([]int, 0)

	for i := 0; i < len(p); i++ {
		freq_p[p[i]-'a']++
	}

	// 整个循环从l==0;r==-1这个空窗口开始
	// 到l==len(s);r==len(s)-1这个空窗口截止
	// 在每次循环里逐渐改变窗口,维护freq,并记录当前窗口
	for r+1 < len(s) { // 窗口的左边界在数组范围内,则循环继续
		r++
		freq_s[s[r]-'a']++
		if r-l+1 > len(p) {
			freq_s[s[l]-'a']--
			l++
		}

		if r-l+1 == len(p) && same(freq_s, freq_p) {
			res = append(res, l)
		}
	}
	return res
}
func same(freq_s, freq_p []int) bool {
	for i := 0; i < 26; i++ {
		if freq_s[i] != freq_p[i] {
			return false
		}
	}
	return true
}
