package LeetCode

import (
	"sort"
	"strconv"
)

// 第一思路: 超出时间限制
func groupAnagrams1(strs []string) [][]string {
	res := [][]string{}
	record := make(map[string]bool, len(strs))
	for _, str := range strs {
		record[str] = true
	}
	for i, str := range strs {
		if record[str] {
			record[str] = false // 已经被使用
			tmp := []string{str}
			for k := i + 1; k < len(strs); k++ {
				if isAnagram(str, strs[k]) {
					record[strs[k]] = false //
					tmp = append(tmp, strs[k])
				}
			}
			res = append(res, tmp)

		}
	}
	return res
}

func isAnagram(a, b string) bool {
	record := make(map[rune]int, 26)
	for _, c := range a {
		record[c]++
	}

	for _, c := range b {
		record[c]--
	}
	for _, v := range record {
		if v != 0 {
			return false
		}
	}
	return true
}

// 字符串排序 Using HashMap Using sorted string as key
// Time Complexity: O(n*klogk) where k is the max length of string in strs
// Space Complexity: O(n*k)
func groupAnagrams2(strs []string) [][]string {
	var res [][]string
	record := make(map[string][]string)
	for _, str := range strs {
		s := []rune(str)            // 把错位词的字符顺序重新排列，那么会得到相同的结果
		sort.Sort(sortRunes(s))     // 以此作为key，将所有错位词都保存到字符串数组中
		val, _ := record[string(s)] // 建立key和字符串数组之间的映射
		val = append(val, str)
		record[string(s)] = val
	}
	for _, v := range record {
		res = append(res, v)
	}
	return res
}

type sortRunes []rune

func (s sortRunes) Len() int           { return len(s) }
func (s sortRunes) Less(i, j int) bool { return s[i] < s[j] }
func (s sortRunes) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// 不用给字符串排序
// Using HashMap Using characters frequency as key
// Time Complexity: O(n*k) where k is the max length of string in strs
// Space Complexity: O(n*k)
func groupAnagrams3(strs []string) [][]string {
	var res [][]string
	m := make(map[string][]string)
	for _, str := range strs {
		key := getKey(str)
		m[key] = append(m[key], str)
	}
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

// 用一个大小为26的int数组来统计每个单词中字符出现的次数，
// 然后将int数组转为一个唯一的字符串，跟字符串数组进行映射，
func getKey(s string) string {
	freq := make([]int, 26)
	for _, c := range s {
		freq[c-'a']++
	}
	key := ""
	for _, n := range freq {
		key += strconv.Itoa(n) + "/"
	}
	return key
}
