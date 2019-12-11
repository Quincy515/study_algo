package Lower_Bound_and_Upper_Bound_in_Binary_Search

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

func generateRandomOrderedArray(n, rangeL, rangeR int) []int {
	var arr []int
	for i := 0; i < n; i++ {
		arr = append(arr, int(rand.Intn(n)%(rangeR-rangeL+1))+rangeL)
	}
	sort.Ints(arr)
	return arr
}
func TestSearch(t *testing.T) {
	rand.Seed(time.Now().Unix())
	n, m := 10, 10
	arr := generateRandomOrderedArray(n, 0, m)
	fmt.Println(arr)

	// 使用简单的线性查找法来验证写的二分查找法
	for i := -1; i <= m+1; i++ {
		if firstGreaterOrEqual(arr, i) != firstGreaterOrEqualLinearSearch(arr, i) {
			panic("线性查找失败!")
		}

		if firstGreaterThan(arr, i) != firstGreaterThanLinearSearch(arr, i) {
			panic("线性查找失败!!")
		}

		if lastLessOrEqual(arr, i) != lastLessOrEqualLinearSearch(arr, i) {
			panic("线性查找失败!!!")
		}

		if lastLessThan(arr, i) != lastLessThanLinearSearch(arr, i) {
			panic("线性查找失败!!!!")
		}
	}
	fmt.Println("test completed.")
}
