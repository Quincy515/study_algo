package leetcode

// 暴力解法 Time: O(log(n)), Space: O(1)
func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	for n%2 == 0 { // n对2取模等于0
		n /= 2 // 不断除以2来更新n
	}
	// 退出循环后只要判断n是否等于1即可
	return n == 1 // 等于1则说明整数是2的幂
}

// 如果n大于0并且n的二进制表示中只要一个比特位是1,那么n就是2的幂
// 2 ^ 3 = 8 二进制表示为 1000
// 2^3-1 = 7 二进制表示为 0111
// n & (n-1) == 0
// 位操作 Time: O(1), Space: O(1)
func isPowerOfTwoO1(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}
