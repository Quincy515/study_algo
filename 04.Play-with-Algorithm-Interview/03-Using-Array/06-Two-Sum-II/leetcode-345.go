package LeetCode

func reverseVowels1(s string) string {
	str := []rune(s)
	i, j := 0, len(str)-1
	for i < j {
		for ; i < j && !isVolume(str[i]); i++ {
		} // 从前往后找到第一个元音字母
		for ; i < j && !isVolume(str[j]); j-- {
		} // 从后往前找到第一个元音字母
		if i < j {
			str[i], str[j] = str[j], str[i]
			i, j = i+1, j-1
		}
	}
	return string(str)
}

func reverseVowels2(s string) string {
	str := []rune(s)
	i := nextVowelIndex(str, 0)
	j := prevVowelIndex(str, len(str)-1)
	for i < j {
		str[i], str[j] = str[j], str[i]
		i = nextVowelIndex(str, i+1)
		j = prevVowelIndex(str, j-1)
	}
	return string(str)
}

func nextVowelIndex(s []rune, index int) int {
	for i := index; i < len(s); i++ {
		if isVolume(s[i]) {
			return i
		}
	}
	return len(s)
}

func prevVowelIndex(s []rune, index int) int {
	for i := index; i >= 0; i-- {
		if isVolume(s[i]) {
			return i
		}
	}
	return -1
}

func isVolume(c rune) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' ||
		c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U'
}
