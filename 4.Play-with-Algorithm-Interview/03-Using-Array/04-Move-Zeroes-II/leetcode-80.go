package main

import "fmt"

func removeDuplicates(nums []int) int {
	k, count := 0, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[k] {
			if count < 2 {
				k++
				count++
				nums[k] = nums[i]
			}
		} else {
			k++
			count = 1
			nums[k] = nums[i]
		}
	}
	return k + 1
}

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(removeDuplicates(nums))

	nums1 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums1))
}
