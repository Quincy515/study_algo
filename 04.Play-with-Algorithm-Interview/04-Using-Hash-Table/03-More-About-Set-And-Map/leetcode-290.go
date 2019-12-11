package LeetCode

import (
	"strings"
)

func wordPattern1(pattern string, str string) bool {
	s := strings.Fields(str)
	freq_s := make(map[int]string)
	for i, word := range s {
		freq_s[i] = word
	}

	freq_p := make(map[int]rune)
	for i, c := range pattern {
		freq_p[i] = c
	}

	if len(freq_s) != len(freq_p) {
		return false
	}

	for k1, v1 := range freq_p {
		for k2, v2 := range freq_p {
			if v1 == v2 {
				if freq_s[k1] != freq_s[k2] {
					return false
				}
			} else {
				if freq_s[k1] == freq_s[k2] {
					return false
				}
			}
		}
	}
	return true
}

// HashMap TimeComplexity: O(n) Space Complexity: O(n)
func wordPattern2(pattern string, str string) bool {
	word := strings.Fields(str)
	if len(pattern) != len(word) {
		return false
	}

	map1 := make(map[rune]string)
	map2 := make(map[string]rune)
	for i := 0; i < len(pattern); i++ {
		if _, ok := map1[rune(pattern[i])]; !ok {
			if _, ok := map2[word[i]]; ok {
				return false
			}
			map1[rune(pattern[i])] = word[i]
			map2[word[i]] = rune(pattern[i])
		} else {
			s := map1[rune(pattern[i])]
			if s != word[i] {
				return false
			}
		}
	}
	return true
}

// HashMap 优化
func wordPattern3(pattern string, str string) bool {
	words := strings.Fields(str)
	if len(pattern) != len(words) {
		return false
	}
	map1 := make(map[byte]string)
	map2 := make(map[string]byte)
	for i := 0; i < len(pattern); i++ {
		v_p, ok_p := map1[pattern[i]]
		v_w, ok_w := map2[words[i]]
		if ok_p && v_p != words[i] || ok_w && v_w != pattern[i] {
			return false
		} else {
			map1[pattern[i]] = words[i]
			map2[words[i]] = pattern[i]
		}
	}
	return true
}
