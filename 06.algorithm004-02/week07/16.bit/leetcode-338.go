package leetcode

// Time: O(n*k), Space: O(1)
func countBits(num int) []int {
	d := make([]int, num+1)     // 定义要返回的结果数组
	for i := 0; i <= num; i++ { // 从0遍历到num执行辅助函数
		d[i] = numberOfOne(i) // 计算每个数的二进制表示中1的个数存储到数组中
	}
	return d // 最后返回数组d即可
}

// Time:O(k),Space: O(1) k表示num的二进制表示中1的个数
func numberOfOne(num int) int {
	count := 0 // 初始化计数器为0
	for num != 0 {
		count++
		num &= (num - 1) // 更新整数num
	}
	return count // 循环结束后返回计数器即可
}

// 动态规划 Time: O(n), Space: O(1)
// d(i)表示整数i的二进制表示中1的个数
// 整数i和i-1按位与得到的结果是，二进制表示中最低位的1消除
// 比如12的二进制表示1100，11的二进制表示为1011，按位与得到的结果是1000
// 1100最低位的1已经被消除，这样一来i的二进制表示中1的个数就会比i&(i-1)中的二进制表示中多出一个1
// 即d(i) = d(i&(i-1))+1有了这个递推式
// 每个整数再计算二进制表示中1的个数时，都可以借助前面整数已经计算出的结果，在O(1)的时间计算出来
// 这种方法总的时间复杂度为O(n),不适用额外的辅助空间Space: O(1)
func countBitsOn(num int) []int {
	d := make([]int, num+1)     // 定义要返回的结果数组
	for i := 1; i <= num; i++ { // 从1遍历到num
		// 在递推式中，整数0要作为最基本情况提供给后面的计算使用，并且d[0]默认已经等于0
		// 于是这里i从1开始
		d[i] = d[i&(i-1)] + 1
	}
	return d // 最后返回数组d即可
}
