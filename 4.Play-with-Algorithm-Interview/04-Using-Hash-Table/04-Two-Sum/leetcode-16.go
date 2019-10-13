package LeetCode

import (
	"sort"
)

func threeSumClosest1(nums []int, target int) int {
	res := 1<<31 - 1
	counter := make(map[int]int)
	for _, num := range nums {
		counter[num]++
	}
	var uniqNums []int
	for key := range counter {
		uniqNums = append(uniqNums, key)
	}
	sort.Ints(uniqNums)

	for i := 0; i < len(uniqNums); i++ {
		if abs(uniqNums[i]*3-target) < abs(res-target) && counter[uniqNums[i]] >= 3 {
			res = uniqNums[i] * 3
		}
		for j := i + 1; j < len(uniqNums); j++ {
			if abs(uniqNums[i]*2+uniqNums[j]-target) < abs(res-target) && counter[uniqNums[i]] >= 2 {
				res = uniqNums[i]*2 + uniqNums[j]
			}
			if abs(uniqNums[i]+uniqNums[j]*2-target) < abs(res-target) && counter[uniqNums[j]] >= 2 {
				res = uniqNums[i] + uniqNums[j]*2
			}
			for k := j + 1; k < len(uniqNums); k++ {
				if abs(uniqNums[i]+uniqNums[j]+uniqNums[k]-target) < abs(res-target) {
					res = uniqNums[i] + uniqNums[j] + uniqNums[k]
				}
			}
		}
	}
	return res
}
func threeSumClosest2(nums []int, target int) int {
	res, difference := 0, 1<<15-1
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if abs(nums[i]+nums[j]+nums[k]-target) < difference {
					difference = abs(nums[i] + nums[j] + nums[k] - target)
					res = nums[i] + nums[j] + nums[k]
				}
			}
		}
	}

	return res
}

func threeSumClosest3(nums []int, target int) int {
	n, res, diff := len(nums), 0, 1<<15-1
	if n > 2 {
		sort.Ints(nums)
		for i := 0; i < n-2; i++ {
			for j, k := i+1, n-1; j < k; {
				sum := nums[i] + nums[j] + nums[k]
				if abs(sum-target) < diff {
					res, diff = sum, abs(sum-target)
				}
				if sum == target {
					return res
				} else if sum > target {
					k--
				} else {
					j++
				}
			}
		}
	}
	return res
}

// Sorting + Two Pointer
// Time complexity: O(nlogn) +O(n^2) Space Complexity: O(1)
func threeSumClosest4(nums []int, target int) int {
	sort.Ints(nums)
	diff := abs(nums[0] + nums[1] + nums[2] - target)
	res := nums[0] + nums[1] + nums[2]

	for i := 0; i < len(nums); i++ {
		l, r := i+1, len(nums)-1
		t := target - nums[i]
		for l < r {
			if nums[l]+nums[r] == t {
				return nums[i] + nums[l] + nums[r]
			} else {
				if abs(nums[l]+nums[r]-t) < diff {
					diff = abs(nums[l] + nums[r] - t)
					res = nums[i] + nums[l] + nums[r]
				}
				if nums[l]+nums[r] < t {
					l++
				} else { // nums[l]+nums[r] > t
					r--
				}
			}
		}
	}
	return res
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -1 * x
}
