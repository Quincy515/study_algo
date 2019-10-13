package LeetCode

func numberOfBoomerangs(points [][]int) int {
	res := 0
	for i := 0; i < len(points); i++ {
		record := make(map[int]int, len(points))
		for j := 0; j < len(points); j++ {
			if j != i {
				record[dis(points[i], points[j])]++
			}
		}

	}
}
