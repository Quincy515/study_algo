package LeetCode

func isHappy(n int) bool {
	record := make(map[int]bool)
	for n != 1 {
		n = happy(n)
		if !record[n] {
			record[n] = true
		} else {
			return false
		}
	}
	return true
}

func happy(n int) int {
	res := 0
	for n != 0 {
		res += (n % 10) * (n % 10)
		n /= 10
	}
	return res
}
