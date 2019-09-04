package main

import (
	"fmt"
	"sort"
)

// 时间复杂度O(n),空间复杂度O(1)
func moveZeroes(nums []int) {
	k := 0 // nums中,[0...k)的元素均为非0元素
	// 遍历到第i个元素后,保证[0...i]中所有非0元素
	// 都按照顺序排序再[0...k)中
	// 同时,[k...i]为0元素
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if i != k { // 如果整个数组全是非0元素
				nums[k], nums[i] = nums[i], nums[k]
				k++
			} else { // i == k避免每个非0元素都自己和自己交换位置
				k++
			}
		}
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}

// 当发现非0元素就在tmp位置覆盖
// 如果它是非 0 的，则它现在已经写入了相应的索引，或者如果它是 0，则稍后将进行处理
func moveZeroes1(nums []int) {
	tmp := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[tmp] = nums[i]
			tmp++
		}
	}

	for i := tmp; i < len(nums); i++ {
		nums[i] = 0
	}
}

// 先排序O(nlogn)，再在下标不为0的位置开始重构数组,使用append为什么数组nums没有改变
func moveZeroes2(nums []int) {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	tmp := 0
	for i, n := range nums {
		if n != 0 {
			tmp = i
			break
		}
	}
	nums = append(nums[tmp:], nums[:tmp]...)
	fmt.Println(nums)
}

// 时间复杂度O(n) 空间复杂度O(n)
func moveZeroes3(nums []int) {
	var nonZerosElements []int

	// 将nums中所有非0元素放入nonZeroElements中
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nonZerosElements = append(nonZerosElements, nums[i])
		}
	}

	// 将nonZerosElements中的所有元素依次放入到nums开始的位置
	for i := 0; i < len(nonZerosElements); i++ {
		nums[i] = nonZerosElements[i]
	}

	// 将nums剩余的位置放置为0
	for i := len(nonZerosElements); i < len(nums); i++ {
		nums[i] = 0
	}
}

// 时间复杂度O(n),空间复杂度O(1)
func moveZeroes4(nums []int) {
	k := 0 // nums中,[0...k)的元素均为非0元素
	// 遍历到第i个元素后,保证[0...i]中所有非0元素
	// 都按照顺序排序再[0...k)中
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[k] = nums[i]
			k++
		}
	}

	// 将nums剩余的位置放置0
	for i := k; i < len(nums); i++ {
		nums[i] = 0
	}
}
