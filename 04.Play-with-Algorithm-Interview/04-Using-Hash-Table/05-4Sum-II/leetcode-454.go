package LeetCode

// Time Complexity: O(n^2) Space Complexity: O(n^2)
func fourSumCount(A []int, B []int, C []int, D []int) int {
	record := make(map[int]int)
	for i := 0; i < len(C); i++ {
		for j := 0; j < len(D); j++ {
			record[C[i]+D[j]]++
		}
	}

	res := 0
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			if val, ok := record[-A[i]-B[j]]; ok {
				res += val
			}
		}
	}
	return res
}

func fourSumCount2(A []int, B []int, C []int, D []int) int {
	hashTable1 := make(map[int]int)
	hashTable2 := make(map[int]int)

	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			hashTable1[A[i]+B[j]]++
		}
	}

	for i := 0; i < len(C); i++ {
		for j := 0; j < len(D); j++ {
			hashTable2[C[i]+D[j]]++
		}
	}

	res := 0
	for key, val := range hashTable1 {
		if v, ok := hashTable2[-key]; ok {
			res += val * v
		}
	}
	return res
}
