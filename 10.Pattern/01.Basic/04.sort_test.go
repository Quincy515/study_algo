package main

import (
	"sort"
	"strconv"
	"testing"
)

func TestSort(t *testing.T) {
	// int 排序
	inits := []int{3, 30, 34, 5, 9}
	sort.Ints(inits)
	t.Log(inits)

	// 字符串排序
	strings := []string{"123", "abv", "1b2c", "上海"}
	sort.Strings(strings)
	t.Log(strings)

	// 自定义排序 - LeetCode179.最大数
	// 如：否则定义一个长度为n的字符串数组
	strs := make([]string, len(inits))
	// 遍历整数数组,把整数转成字符串,并存储到字符串数组
	for i := range inits {
		strs[i] = strconv.Itoa(inits[i])
	}
	sort.Slice(strs, func(i, j int) bool {
		// 先把两种字符串的拼接结果写出来
		s12 := strs[i] + strs[j]
		s21 := strs[j] + strs[i]
		// 如果字符串s12大于s21，那么我们希望s1排在s2前面。
		return s12 > s21
	})
	t.Log(strs)
}
