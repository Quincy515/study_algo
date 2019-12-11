package LeetCode

import (
	"strings"
)

func isPalindrome1(s string) bool {
	s = strings.ToLower(s) // 统一转为小写
	str := []rune(s)
	i, j := 0, len(str)-1
	for i < j {
		// 跳过非法字符
		for !isAlphanumeric(str[i]) {
			i++
			if i == len(str) {
				return true
			}
		}
		for !isAlphanumeric(str[j]) {
			j--
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func isAlphanumeric(c rune) bool {
	if 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || '0' <= c && c <= '9' {
		return true
	}
	return false
}

func isPalindrome2(s string) bool {
	charMap := make([]rune, 256)
	for i := 0; i < 10; i++ { // 映射 '0' 到 '9'
		charMap[i+'0'] = rune(1 + i) // numeric
	}
	// 映射 'a' 到 'z' 和映射 'A' 到 'Z'
	for i := 0; i < 26; i++ {
		charMap[i+'a'] = rune(11 + i)
		charMap[i+'A'] = rune(11 + i)
	}

	pChars := []rune(s)
	start, end := 0, len(pChars)-1
	for start < end {
		// 得到映射后的数字
		cS := charMap[pChars[start]]
		cE := charMap[pChars[end]]
		if cS != 0 && cE != 0 {
			if cS != cE {
				return false
			}
			start++
			end--
		} else {
			if cS == 0 {
				start++
			}
			if cE == 0 {
				end--
			}
		}
	}
	return true
}

func isPalindrome3(s string) bool {
	str := []rune(s)
	i := next_alpha_numeric(str, 0)
	j := prev_alpha_numeric(str, len(str)-1)
	for i <= j {
		if strings.ToLower(string(str[i])) != strings.ToLower(string(str[j])) {
			return false
		}
		i = next_alpha_numeric(str, i+1)
		j = prev_alpha_numeric(str, j-1)
	}
	return true
}

func next_alpha_numeric(s []rune, index int) int {
	for i := index; i < len(s); i++ {
		if isAlphanumeric(s[i]) {
			return i
		}
	}
	return len(s)
}

func prev_alpha_numeric(s []rune, index int) int {
	for i := index; i >= 0; i-- {
		if isAlphanumeric(s[i]) {
			return i
		}
	}
	return -1
}
