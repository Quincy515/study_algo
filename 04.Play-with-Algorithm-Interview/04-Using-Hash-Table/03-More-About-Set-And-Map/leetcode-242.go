package LeetCode

func isAnagram(s string, t string) bool {
	record := make(map[rune]int)
	for _, c := range s {
		record[c]++
	}

	for _, c := range t {
		if _, ok := record[c]; ok {
			record[c]--
		} else {
			return false
		}
	}

	for _, v := range record {
		if v != 0 {
			return false
		}
	}
	return true
}
