package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	// 数据规模每次增大10倍进行测试,也可以试验一下数据规模每次增大2倍
	for x := 1; x <= 9; x++ {
		n := int(math.Pow10(x))
		startTime := time.Now()
		sum := 0
		for i := 0; i < n; i++ {
			sum += i
		}
		spent := time.Now().Sub(startTime)
		fmt.Println("sum = ", sum)
		fmt.Println("10^", x, " : ", spent)
	}
}
