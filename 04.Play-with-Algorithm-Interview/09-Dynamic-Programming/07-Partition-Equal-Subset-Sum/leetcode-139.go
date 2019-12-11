package leetcode

// Time Complexity: O(n^2), Space Complexity: O(n+m)
func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	d := make([]bool, n+1)
	d[0] = true
	set := make(map[string]bool)
	for _, s := range wordDict {
		set[s] = true
	}
	for i := 1; i <= n; i++ {
		for j := i - 1; j >= 0; j-- {
			ok, _ := set[s[j:i]]
			if d[j] && ok {
				d[i] = true
				break
			}
		}
	}
	return d[n]
}
