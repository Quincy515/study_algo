package LeetCode

func minWindow1(s string, t string) string {
	freq_t := make([]int, 256)
	for i := 0; i < len(t); i++ {
		freq_t[t[i]]++
	}

	freq_s := make([]int, 256)
	sCnt := 0
	minLength := len(s) + 1
	startIndex := -1

	l, r := 0, -1
	for l < len(s) {
		if r+1 < len(s) && sCnt < len(t) {
			freq_s[s[r+1]]++
			if freq_s[s[r+1]] <= freq_t[s[r+1]] {
				sCnt++
			}
			r++
		} else {
			if sCnt == len(t) && r-l+1 < minLength {
				minLength = r - l + 1
				startIndex = l
			}
			freq_s[s[l]]--
			if freq_s[s[l]] < freq_t[s[l]] {
				sCnt--
			}
			l++
		}
	}
	if startIndex != -1 {
		return s[startIndex : startIndex+minLength]
	}
	return ""
}
