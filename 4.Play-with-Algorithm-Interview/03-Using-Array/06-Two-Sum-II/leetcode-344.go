package LeetCode

import "fmt"

func reverseString1(s []byte) {
	i, j := 0, len(s)-1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
	fmt.Println(string(s))
}

func reverseString2(s []byte) {
	size := len(s) - 1
	for i := 0; i < len(s)/2; i++ {
		s[i], s[size-i] = s[size-i], s[i]
	}
	fmt.Println(string(s))
}

func reverseString3(s []byte) {
	for i, j := 0, len(s)-1; i < j; {
		s[i], s[j] = s[j], s[i]
		i, j = i+1, j-1
	}
	fmt.Println(string(s))
}
