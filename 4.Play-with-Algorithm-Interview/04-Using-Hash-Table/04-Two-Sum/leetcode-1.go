package LeetCode

func twoSum1(nums []int, target int) []int {
	record := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		record[i] = target - nums[i]
	}
	for i, num := range nums {
		for k, v := range record {
			if num == v && k != i {
				return []int{i, k}
			}
			//delete(record, k)
		}
	}
	return []int{}
}

// 查找表
func twoSum2(nums []int, target int) []int {
	record := make(map[int]int)
	for i, n := range nums {
		record[target-n] = i
	}

	for i, n := range nums {
		if v, ok := record[n]; ok && v != i {
			return []int{i, v}
		}
	}
	return []int{}
}

// 优化查找表,哈希表
func twoSum3(nums []int, target int) []int {
	m := make(map[int]int)
	result := []int{}
	// 循环遍历元素,使用map来存储元素和下标之间的映射关系
	for i, num := range nums {
		// 判断map的key是否存在
		if j, ok := m[target-num]; ok {
			result = append(result, j)
			result = append(result, i)
		}
		m[num] = i // 使用map来存储元素和下标之间的映射关系
	}
	return result
}
