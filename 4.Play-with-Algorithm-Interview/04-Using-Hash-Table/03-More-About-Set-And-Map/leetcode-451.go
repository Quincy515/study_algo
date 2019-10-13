package LeetCode

import (
	"sort"
)

func frequencySort1(s string) string {
	record := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		record[s[i]]++
	}
	var tmp []byte
	for len(record) > 0 {
		k, v := maxFrequency(record)
		for i := 0; i < v; i++ {
			tmp = append(tmp, k)
		}
		delete(record, k)
	}
	return string(tmp)
}

func maxFrequency(record map[byte]int) (key byte, val int) {
	for k, v := range record {
		if v > val {
			val = v
			key = k
		}
	}
	return
}

func frequencySort2(s string) string {
	freq := make(map[rune]int, len(s))
	for _, c := range s {
		freq[c]++
	}
	ss := make([]string, 0, len(freq))
	for k, v := range freq {
		tmp := make([]rune, v)
		for i := range tmp {
			tmp[i] = k
		}
		ss = append(ss, string(tmp))
	}
	sort.Sort(str(ss))
	res := ""
	for _, v := range ss {
		res += v
	}
	return res
}

type str []string

func (s str) Len() int           { return len(s) }
func (s str) Less(i, j int) bool { return len(s[i]) > len(s[j]) }
func (s str) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
