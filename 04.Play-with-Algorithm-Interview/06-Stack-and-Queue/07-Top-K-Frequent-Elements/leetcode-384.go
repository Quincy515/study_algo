package leetcode

import "math/rand"

type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	return Solution{nums}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.nums
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	nums := make([]int, len(this.nums))
	buf := make([]int, len(this.nums)) // 创建新数组
	copy(buf, this.nums)
	for i := range nums { // 循环从原数组中随机抽取元素填入新数组
		j := rand.Intn(len(buf))
		nums[i] = buf[j] // 并将已抽取的元素从原数组中删除
		buf = append(buf[0:j], buf[j+1:]...)
	}
	return nums
}

func (this *Solution) Shuffle1() []int {
	nums := make([]int, len(this.nums))
	copy(nums, this.nums)
	// 直接调用库函数打乱数组顺序
	rand.Shuffle(len(nums), func(i int, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})
	return nums
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
