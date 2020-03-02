package leetcode

import "sort"

// 辅助递归函数，输入是正整数数组和目标值，还包含指向数组的开始下标start，以及当前数字组合elem和最后结果集合result
func combSum(nums []int, target, start int, elem []int, result *[][]int) {
	if target == 0 { // 如果目标值已经减到0
		tmp := make([]int, len(elem)) // 说明当前组合elem是一个有效的组合
		copy(tmp, elem)
		*result = append(*result, tmp) // 把它加入结果集合中
		return                         // 然后返回
	}
	if target < 0 { // 如果目标值已经减到小于0
		return // 说明当前组合elem是一个无效组合,直接返会
	} // 如果不是上面两种情况，则说明目标值仍然大于0，
	// 继续在数组中取值相加,从下标start开始继续遍历数组
	for i := start; i < len(nums); i++ {
		elem = append(elem, nums[i]) // 取出一个数字加入到当前组合
		// 然后递归调用自己，目标值要减去当前数字,当前数字的下标是i,
		// 由于我们可以从后取数，因此递归下一轮的start从i开始
		combSum(nums, target-nums[i], i, elem, result)
		elem = elem[:len(elem)-1] // 退递归后需要把当前数字组合最后加入的数字移除
	}
}

// Time: O(n^(target/min)), Space: O(target/min) n是数组长度，min是数组中最小值
func combinationSum(candidates []int, target int) [][]int {
	var result [][]int // 定义结果集合result
	var elem []int     // 定义当前数字组合elem
	// 调用辅助递归函数，开始下标start等于0，最后返回结果即可
	combSum(candidates, target, 0, elem, &result)
	return result
}

// 剪枝优化，先对数组进行排序，这样只要当遇到的某个数字使得组合加起来大于目标值，那它后面的数字就都可以不看了，直接回溯到上一个位置，
// 优化只需要再上面的代码上加两行，一是先对数组排序，然后在辅助递归函数的for循环中，先检查当前数字是否大于目标值，如果是直接跳出循环
// 因为数组有序，因此后面的数字也肯定大于目标值，这个优化可以减少操作次数，但时间复杂度在数量级上和优化前一样
// 辅助递归函数，输入是正整数数组和目标值，还包含指向数组的开始下标start，以及当前数字组合elem和最后结果集合result
func combSum(nums []int, target, start int, elem []int, result *[][]int) {
	if target == 0 { // 如果目标值已经减到0
		tmp := make([]int, len(elem)) // 说明当前组合elem是一个有效的组合
		copy(tmp, elem)
		*result = append(*result, tmp) // 把它加入结果集合中
		return                         // 然后返回
	}
	if target < 0 { // 如果目标值已经减到小于0
		return // 说明当前组合elem是一个无效组合,直接返会
	} // 如果不是上面两种情况，则说明目标值仍然大于0，
	// 继续在数组中取值相加,从下标start开始继续遍历数组
	for i := start; i < len(nums); i++ {
		if nums[i] > target {
			break
		}
		elem = append(elem, nums[i]) // 取出一个数字加入到当前组合
		// 然后递归调用自己，目标值要减去当前数字,当前数字的下标是i,
		// 由于我们可以从后取数，因此递归下一轮的start从i开始
		combSum(nums, target-nums[i], i, elem, result)
		elem = elem[:len(elem)-1] // 退递归后需要把当前数字组合最后加入的数字移除
	}
}

// Time: O(n^(target/min)), Space: O(target/min) n是数组长度，min是数组中最小值
func combinationSum(candidates []int, target int) [][]int {
	var result [][]int // 定义结果集合result
	var elem []int     // 定义当前数字组合elem
	sort.Ints(candidates)
	// 调用辅助递归函数，开始下标start等于0，最后返回结果即可
	combSum(candidates, target, 0, elem, &result)
	return result
}
