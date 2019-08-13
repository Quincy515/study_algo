package Selection_Sort

import (
	"math/rand"
	"time"
)

// 生成n个元素的随机数组，每个元素的随机范围为[rangeL, rangeR]
func GenerateRandomArray(n, rangeL, rangeR int) []int {
	if rangeL > rangeR {
		panic("rangeL cannot be bigger than rangeR!")
	}

	var arr []int
	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		arr = append(arr, rand.Int()%(rangeR-rangeL+1)+rangeL)
	}
	return arr
}
