package leetcode

import (
	"bytes"
	"fmt"
)

func maxPoints(points [][]int) int {
	if len(points) < 2 {
		return len(points)
	}
	res := 1                           // 记录直线上最多的点数
	for i := 0; i < len(points); i++ { // 遍历枢纽点i
		// 对于每一个points[i]设立一个查找表，key是斜率，value是相同斜率的点个数
		record := make(map[string]int)
		samePoint := 0                     // 相同的点的个数
		for j := 0; j < len(points); j++ { // 遍历其余所有的点
			if points[i][0] == points[j][0] && points[i][1] == points[j][1] { // 索引j和i相等，即找到一个相同的点
				samePoint++
			} else { // 索引j和i对应的值不相等，即找到一个不相同的点
				record[getPairStr(slope(points[i], points[j]))]++ // 记录这个斜率出现的次数
			}
		}
		// 如果记录为空并且所有点都在同一点
		res = max(res, samePoint)
		for _, v := range record {
			res = max(res, v+samePoint)
		}
	}
	return res
}

// 辅助函数 计算两点之间的斜率
func slope(pa, pb []int) []int {
	dy := pa[1] - pb[1]
	dx := pa[0] - pb[0]
	if dx == 0 {
		return []int{1, 0}
	}
	if dy == 0 {
		return []int{0, 1}
	}
	g := gcd(abs(dy), abs(dx))
	dy /= g
	dx /= g
	if dx < 0 {
		dy = -dy
		dx = -dx
	}
	return []int{dy, dx}
}

func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	if a%b == 0 {
		return b
	}
	return gcd(b, a%b)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func getPairStr(p []int) string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("%d/%d\n", p[0], p[1]))
	return buffer.String()
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
