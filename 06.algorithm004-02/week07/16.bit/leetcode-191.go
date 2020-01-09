package leetcode

// Time: O(m),Space: O(1) m表示num的二进制位数
func hammingWeight(num uint32) int {
	mask, count := uint32(1), 0 // 初始化掩码为1,计数器为0
	for mask != 0 {             // 当掩码不为0时,执行以下操作
		// 整数与掩码按位与
		if (num & mask) != 0 { // 不等于0
			count++ // 计数器加一
		}
		mask <<= 1 // 掩码向左移动一位
	}
	return count // 循环结束后,返回计数器即可
}

// Time:O(k),Space: O(1) k表示num的二进制表示中1的个数
func hammingWeightOk(num uint32) int {
	count := 0 // 初始化计数器为0
	for num != 0 {
		count++
		num &= (num - 1) // 更新整数num
	}
	return count // 循环结束后返回计数器即可
}
