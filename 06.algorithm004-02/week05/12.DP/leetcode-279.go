package leetcode

import "math"

// 动态规划 Time: O(n ^ 3/2), Space: O(n)
func numSquares(n int) int {
	d := make([]int, n+1)     // 定义大小为n+1的辅助数组d用于表示状态
	d[0] = 0                  // 初始化边界d[0]
	for i := 1; i <= n; i++ { // i从1遍历到n
		d[i] = i                    // 初始化d[i]=i，因为整数i可以分解成的完全平方数的数量不会超过i，即i和1相加
		for j := 1; j*j <= i; j++ { // 接着穷举j，约束条件是j的平方小于等于i
			d[i] = min(d[i], d[i-j*j]+1)
		}
	}
	return d[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 数学解法 Time:O(n^1/2), Space:O(1)）
func numSquaresMath(n int) int {
	if isSquare(n) { // 首先判断整数n是否是完全平方数
		return 1 // 如果是就返回1
	}
	for i := 1; i*i <= n; i++ {
		if isSquare(n - i*i) {
			return 2
		}
	}
	for n%4 == 0 {
		n /= 4
	}
	if n%8 != 7 {
		return 3
	}
	return 4
}

// 辅助函数 用于判断整数n是否是一个完全平方数
func isSquare(n int) bool {
	a := int(math.Sqrt(float64(n))) // 对整数n开根号并取整赋值给变量a
	return a*a == n                 // 然后判断a*a是否等于n即可
}
