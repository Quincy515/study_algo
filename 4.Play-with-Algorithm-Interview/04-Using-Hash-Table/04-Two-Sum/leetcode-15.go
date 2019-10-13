package LeetCode

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	map_t := make(map[int]bool, 10)
	for i := 0; i < len(nums); i++ {
		map_t[nums[i]] = true
	}
	var tmp []int
	for k, _ := range map_t {
		tmp = append(tmp, k)
	}
	nums = tmp

	m := make(map[int]bool)
	result := [][]int{}
	// 循环遍历元素,使用map来存储元素和下标之间的映射关系
	for i, num := range nums {
		var res []int
		for j, n := range nums {
			// 判断map的key是否存在
			if _, ok := m[-n-num]; ok && j > i {
				res = []int{-n - num, num, n}
			}
		}
		m[num] = true // 使用map来存储元素和下标之间的映射关系
		if len(res) != 0 && !zero(res) {
			result = append(result, res)
		} else if len(res) != 0 && zero(res) {
			return append(result, []int{0, 0, 0})
		}
	}
	return result
}

func zero(a []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != 0 {
			return false
		}
	}
	return true
}

//
//func duplicate(a [][]int, b []int) bool {
//	for i := 0; i < len(a); i++ {
//		for j := 0; j < len(a[i]); j++ {
//			for _, v2 := range b {
//				if a[i][j] == v2 {
//					v2 = 0
//				}
//			}
//		}
//	}
//	for i := 0; i < len(b); i++ {
//		if b[i] != 0 {
//			return false
//		}
//	}
//
//	return true
//}

// Using Hash Table to store all the numbers
// Time Complexity: O(n^2) Space Complexity: O(n)
func threeSum1(nums []int) [][]int {
	counter := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		counter[nums[i]]++
	}
	var res [][]int
	if counter[0] >= 3 {
		res = append(res, []int{0, 0, 0})
	}
	// Remove duplications
	var tmp []int
	for k := range counter {
		tmp = append(tmp, k)
	}
	nums = tmp
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]*2+nums[j] == 0 && counter[nums[i]] >= 2 {
				res = append(res, []int{nums[i], nums[i], nums[j]})
			}
			if nums[i]+nums[j]*2 == 0 && counter[nums[j]] >= 2 {
				res = append(res, []int{nums[i], nums[j], nums[j]})
			}
			c := 0 - nums[i] - nums[j]
			if c > nums[j] && counter[c] != 0 {
				res = append(res, []int{nums[i], nums[j], c})
			}
		}
	}
	return res
}

// Using two pointers technique
// Time Complexity: O(n^2) Space Complexity: O(n)
func threeSum2(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	index := 0
	for index < len(nums) {
		if nums[index] > 0 {
			break
		}
		bindex, cindex := index+1, len(nums)-1
		for bindex < cindex {
			if nums[bindex]+nums[cindex] == -nums[index] {
				res = append(res, []int{nums[index], nums[bindex], nums[cindex]})
				// continue to look for other pairs
				bindex = nex_num_index(nums, bindex)
				cindex = pre_num_index(nums, cindex)
			} else if nums[bindex]+nums[cindex] < -nums[index] {
				bindex = nex_num_index(nums, bindex)
			} else { // nums[bindex] + nums[cindex] > -nums[index]
				cindex = pre_num_index(nums, cindex)
			}
		}
		index = nex_num_index(nums, index)
	}
	return res
}

func nex_num_index(nums []int, cur int) int {
	for i := cur + 1; i < len(nums); i++ {
		if nums[i] != nums[cur] {
			return i
		}
	}
	return len(nums)
}

func pre_num_index(nums []int, cur int) int {
	for i := cur - 1; i >= 0; i-- {
		if nums[i] != nums[cur] {
			return i
		}
	}
	return -1
}

func threeSum3(nums []int) [][]int {
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
		if uniqNums[i]*3 == 0 && counter[uniqNums[i]] >= 3 {
			res = append(res, []int{0, 0, 0})
		}
		for j := i + 1; j < len(uniqNums); j++ {
			if uniqNums[i]*2+uniqNums[j] == 0 && counter[uniqNums[i]] >= 2 {
				res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[j]})
			}
			if uniqNums[i]+uniqNums[j]*2 == 0 && counter[uniqNums[j]] >= 2 {
				res = append(res, []int{uniqNums[i], uniqNums[j], uniqNums[j]})
			}
			c := 0 - uniqNums[i] - uniqNums[j]
			if c > uniqNums[j] && counter[c] > 0 {
				res = append(res, []int{uniqNums[i], uniqNums[j], c})
			}
		}
	}
	return res
}
