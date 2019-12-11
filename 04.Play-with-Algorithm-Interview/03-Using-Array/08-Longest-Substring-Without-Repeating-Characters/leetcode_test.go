package LeetCode

import (
	"fmt"
	"testing"
)

func TestLeetCode3(t *testing.T) {
	s1 := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring1(s1))
	fmt.Println(lengthOfLongestSubstring2(s1))
	fmt.Println(lengthOfLongestSubstring3(s1))
	fmt.Println(lengthOfLongestSubstring4(s1))
	fmt.Println(lengthOfLongestSubstring5(s1))
}

func TestLeetCode438(t *testing.T) {
	s1, p1 := "cbaebabacd", "abc"
	fmt.Println(findAnagrams1(s1, p1))
	s2, p2 := "abab", "ab"
	fmt.Println(findAnagrams1(s2, p2))
}

func TestLeetCode76(t *testing.T) {
	s1, t1 := "ADOBECODEBANC", "ABC"
	fmt.Println(minWindow1(s1, t1))

	s2, t2 := "aaflslflsldkalskaaa", "aaa"
	fmt.Println(minWindow1(s2, t2))
}
