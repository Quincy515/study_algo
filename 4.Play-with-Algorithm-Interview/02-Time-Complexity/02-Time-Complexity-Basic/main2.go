package main

import "fmt"

func sum1(n int) int {
	assert(n >= 0)
	ret := 0
	for i := 0; i <= n; i++ {
		ret += i
	}
	return ret
}

func sum2(n int) int {
	assert(n >= 0)
	if n == 0 {
		return 0
	}
	return n + sum2(n-1)
}

func main() {
	fmt.Println("sum1(10000) = ", sum1(10000))
	fmt.Println("sum2(10000) = ", sum2(10000))
}

func assert(exp bool) {
	if !exp {
		panic("断言失败")
	}
}
