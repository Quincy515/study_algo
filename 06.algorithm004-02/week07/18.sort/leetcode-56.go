package leetcode

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	// 按区间开始下标从小到大排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var result [][]int // 定义结果list
	result = append(result, intervals[0])
	for i := 1; i < len(intervals); i++ { // 遍历区间数组
		var a = len(result) - 1 // 先把结果数组大小拿出来，方便取数组最后一个元素
		var b = len(result[a]) - 1
		// 结果数组最后的区间，和当前区间有重叠，更新这个区间的结束下标，两个区间结束下标中较大值
		if result[a][b] >= intervals[i][0] {
			var j = len(intervals[i]) - 1
			result[a][b] = max(result[a][b], intervals[i][j])
		} else { // 结果数组为空或者结果数组最后区间结束下标小于当前区间开始下标
			result = append(result, intervals[i]) // 当前区间加入结果数组

		}
	}
	return result // 最后返回结果数组即可
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
