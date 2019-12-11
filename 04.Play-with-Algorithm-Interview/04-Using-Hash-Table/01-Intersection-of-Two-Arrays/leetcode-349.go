package LeetCode

func intersection1(nums1 []int, nums2 []int) []int {
	var res []int
	uniqueNums1 := make(map[int]bool)
	for _, num := range nums1 {
		uniqueNums1[num] = true
	}

	for _, num := range nums2 {
		if uniqueNums1[num] == true {
			res = append(res, num)
			uniqueNums1[num] = false
		}
	}
	return res
}

func intersection2(nums1 []int, nums2 []int) []int {
	record := make(map[int]bool) // 记录nums1中出现的元素
	for i := 0; i < len(nums1); i++ {
		record[nums1[i]] = true // 使用map不能承载重复元素
	}

	resultSet := make(map[int]bool)
	for i := 0; i < len(nums2); i++ { // 遍历nums2中的所有元素
		if record[nums2[i]] == true { // 查看在nums2中的元素在nums1中是否存在
			resultSet[nums2[i]] = true // 记录元素,每个元素记录一次,所以使用set
		}
	}

	var resultVector []int
	for k := range resultSet { // 将resultSet中所有元素保存到vector中
		resultVector = append(resultVector, k)
	}
	return resultVector
}
