package LeetCode

func singleNumber1(nums []int) int {
	ans := 0
	for _, num := range nums {
		ans ^= num
	}
	return ans
}

func singleNumber2(nums []int) int {
	record := make(map[int]int)
	for _, num := range nums {
		record[num]++
	}

	for k, v := range record {
		if v == 1 {
			return k
		}
	}
	return 0
}
