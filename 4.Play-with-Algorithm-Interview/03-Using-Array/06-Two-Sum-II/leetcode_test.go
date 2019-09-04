package LeetCode

import (
	"fmt"
	"testing"
)

func TestLeetCode167(t *testing.T) {
	numbers, target := []int{2, 7, 11, 15}, 9
	fmt.Println(twoSum1(numbers, target))
	fmt.Println(twoSum2(numbers, target))
	fmt.Println()
}

func TestLeetCode125(t *testing.T) {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome1(s))
	fmt.Println(isPalindrome2(s))
	fmt.Println(isPalindrome3(s))

	s1 := "race a car"
	fmt.Println(isPalindrome1(s1))
	fmt.Println(isPalindrome2(s1))
	fmt.Println(isPalindrome3(s1))

	fmt.Println()
}

func TestLeetCode344(t *testing.T) {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString1(s)
	reverseString2(s)
	reverseString3(s)
	fmt.Println()
}

func TestLeetCode345(t *testing.T) {
	s := "hello"
	reverseVowels1(s)
	reverseVowels2(s)
}

func TestLeetCode11(t *testing.T) {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height)) // 49
	height1 := []int{1, 1}
	fmt.Println(maxArea(height1)) // 1
	height2 := []int{0, 2}
	fmt.Println(maxArea(height2)) // 0
	height3 := []int{2, 3, 10, 5, 7, 8, 9}
	fmt.Println(maxArea(height3)) // 36
}
