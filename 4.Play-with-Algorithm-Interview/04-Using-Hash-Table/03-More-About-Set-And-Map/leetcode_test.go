package LeetCode

import (
	"fmt"
	"testing"
)

func TestLeetCode136(t *testing.T) {
	nums := []int{2, 2, 1}
	fmt.Println(singleNumber1(nums))
	fmt.Println(singleNumber2(nums))
}

func TestLeetCode242(t *testing.T) {
	s, tt := "anagram", "nagaram"
	s1, tt1 := "rat", "car"
	fmt.Println(isAnagram(s, tt))
	fmt.Println(isAnagram(s1, tt1))
}
func TestLeetCode202(t *testing.T) {
	fmt.Println(isHappy(19))
	fmt.Println(isHappy(4))
}

func TestLeetCode290(t *testing.T) {
	pattern1, str1 := "abba", "dog cat cat dog"
	pattern2, str2 := "abba", "dog cat cat fish"
	pattern3, str3 := "aaaa", "dog cat cat dog"
	pattern4, str4 := "abba", "dog dog dog dog"
	pattern5, str5 := "aaa", "aa aa aa aa"
	fmt.Println(wordPattern2(pattern1, str1))
	fmt.Println(wordPattern2(pattern2, str2))
	fmt.Println(wordPattern2(pattern3, str3))
	fmt.Println(wordPattern3(pattern4, str4))
	fmt.Println(wordPattern3(pattern5, str5))
}

func TestLeetCode205(t *testing.T) {
	s1, t1 := "egg", "add"
	s2, t2 := "foo", "bar"
	s3, t3 := "paper", "title"
	s4, t4 := "ab", "aa"
	fmt.Println(isIsomorphic1(s1, t1))
	fmt.Println(isIsomorphic1(s2, t2))
	fmt.Println(isIsomorphic2(s3, t3))
	fmt.Println(isIsomorphic2(s4, t4))
}

func TestLeetCode451(t *testing.T) {
	s1 := "tree"
	s2 := "cccaaa"
	s3 := "Aabb"
	fmt.Println(frequencySort1(s1))
	fmt.Println(frequencySort2(s2))
	fmt.Println(frequencySort2(s3))
}
