package LeetCode

import "sort"

// Using HashTable
// Sort the array list. Store every different c + d == t first
// For  every different number a and b, try to find a pair (c,d), which a+b+c+d==0
// Time Complexity: O(nlogn)+O(n^2)+O(n^3) Space Complexity: O(n^2)
func fourSum1(nums []int, target int) [][]int {
	var res [][]int
	counter := make(map[int]int)
	for _, value := range nums {
		counter[value]++
	}

	var uniqNums []int
	for key := range counter {
		uniqNums = append(uniqNums, key)
	}
	sort.Ints(uniqNums)

	for i := 0; i < len(uniqNums); i++ {
		if uniqNums[i]*4 == target && counter[uniqNums[i]] >= 4 {
			res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[i], uniqNums[i]})
		}
		for j := i + 1; j < len(uniqNums); j++ {
			if uniqNums[i]*2+uniqNums[j]*2 == target && counter[uniqNums[i]] >= 2 && counter[uniqNums[j]] >= 2 {
				res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[j], uniqNums[j]})
			}
			if uniqNums[i]*3+uniqNums[j] == target && counter[uniqNums[i]] >= 3 {
				res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[i], uniqNums[j]})
			}
			if uniqNums[i]+uniqNums[j]*3 == target && counter[uniqNums[j]] >= 3 {
				res = append(res, []int{uniqNums[i], uniqNums[j], uniqNums[j], uniqNums[j]})
			}
			for k := j + 1; k < len(uniqNums); k++ {
				if uniqNums[i]*2+uniqNums[j]+uniqNums[k] == target && counter[uniqNums[i]] >= 2 {
					res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[j], uniqNums[k]})
				}
				if uniqNums[i]+uniqNums[j]*2+uniqNums[k] == target && counter[uniqNums[j]] >= 2 {
					res = append(res, []int{uniqNums[i], uniqNums[j], uniqNums[j], uniqNums[k]})
				}
				if uniqNums[i]+uniqNums[j]+uniqNums[k]*2 == target && counter[uniqNums[k]] >= 2 {
					res = append(res, []int{uniqNums[i], uniqNums[j], uniqNums[k], uniqNums[k]})
				}
				c := target - uniqNums[i] - uniqNums[j] - uniqNums[k]
				if c > uniqNums[k] && counter[c] > 0 {
					res = append(res, []int{uniqNums[i], uniqNums[j], uniqNums[k], c})
				}
			}
		}
	}
	return res
}

// Two Pointer
// Sort the array first. For every different number a and b, try to find pair(c,d).which a+b+c+d=0.
// Using this way, we dont need to see whether the triplet is a repeated one
// Time Complexity: O(nlogn) + O(n^3) Space Complexity: O(1)
func fourSum2(nums []int, target int) [][]int {
	var res [][]int
	n := len(nums)
	sort.Ints(nums)
	for i := 0; i <= n-4; i = nextNumberIndex(nums, i) {
		for j := i + 1; j <= n-3; j = nextNumberIndex(nums, j) {
			t := target - nums[i] - nums[j]
			if nums[j+1]+nums[j+2] > t || nums[n-1]+nums[n-2] < t {
				continue
			}
			p1 := j + 1
			p2 := len(nums) - 1
			if p1 >= p2 {
				break
			}

			for p1 < p2 {
				if nums[p1]+nums[p2] == t {
					res = append(res, []int{nums[i], nums[j], nums[p1], nums[p2]})
					p1 = nextNumberIndex(nums, p1)
					p2 = prevNumberIndex(nums, p2)
				} else if nums[p1]+nums[p2] < t {
					p1 = nextNumberIndex(nums, p1)
				} else { // nums[p1]+nums[p2]>t
					p2 = prevNumberIndex(nums, p2)
				}
			}
		}
	}
	return res
}

func nextNumberIndex(nums []int, index int) int {
	for i := index + 1; i < len(nums); i++ {
		if nums[i] != nums[index] {
			return i
		}
	}
	return len(nums)
}

func prevNumberIndex(nums []int, index int) int {
	for i := index - 1; i >= 0; i-- {
		if nums[i] != nums[index] {
			return i
		}
	}
	return -1
}
