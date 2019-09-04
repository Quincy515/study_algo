package LeetCode

func sortColors1(nums []int) {
	mark := [...]int{0, 1, 2} // 数组
	for i, _ := range mark {
		mark[i] = 0
	}

	for _, v := range nums {
		mark[v] += 1
	}
	k := 0
	for i, v := range mark {
		for j := 0; j < v; j++ {
			nums[k] = i
			k++
		}
	}
}

//优化代码
func sortColors2(nums []int) {
	mark := [3]int{0} // 数组
	for _, v := range nums {
		mark[v] += 1
	}
	k := 0
	for i, v := range mark {
		for j := 0; j < v; j++ {
			nums[k] = i
			k++
		}
	}
}

// 计数排序的思路,对整个数组遍历了两遍
// 时间复杂度: O(n) 空间复杂度: O(k), k为元素的取值范围
func sortColors3(nums []int) {
	count := [3]int{0} // 存放0,1,2三个元素的频率
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 || nums[i] > 2 {
			panic("错误处理!")
		}
		count[nums[i]]++
	}

	index := 0
	for i := 0; i < count[0]; i++ {
		nums[index] = 0
		index++
	}
	for i := 0; i < count[1]; i++ {
		nums[index] = 1
		index++
	}
	for i := 0; i < count[2]; i++ {
		nums[index] = 2
		index++
	}
}

// 时间复杂度度: O(n),空间复杂度: O(1),只遍历一遍数组
func sortColors4(nums []int) {
	zero := -1       // nums[0...zero] == 0
	two := len(nums) // nums[two...n-1] == 2
	for i := 0; i < two; {
		if nums[i] == 1 {
			i++
		} else if nums[i] == 2 {
			two--
			nums[i], nums[two] = nums[two], nums[i]
		} else { // nums[i] == 0
			zero++
			nums[zero], nums[i] = nums[i], nums[zero]
			i++
		}
	}
}
