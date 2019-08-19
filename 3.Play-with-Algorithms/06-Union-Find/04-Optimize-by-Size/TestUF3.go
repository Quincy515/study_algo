package Optimize_by_Size

import (
	"fmt"
	"math/rand"
	"time"
)

// 测试并查集的辅助函数
// 测试第一版的并查集,测试元素个数为n
func TestUF3(n int) {
	rand.Seed(time.Now().Unix())
	uf := NewUnionFind(n)
	startTime := time.Now()

	// 进行n次操作,每次随机选择两个元素进行合并操作
	for i := 0; i < n; i++ {
		a, b := rand.Int()%n, rand.Int()%n
		uf.UnionElements(a, b)
	}
	// 再进行n次操作,每次随机选择两个元素,查询他们是否同属于一个集合
	for i := 0; i < n; i++ {
		a, b := rand.Int()%n, rand.Int()%n
		uf.IsConnected(a, b)
	}
	spentTime := time.Now().Sub(startTime)

	// 打印输出对这2n个操作的耗时
	fmt.Println("UF3, ", 2*n, " ops, ", spentTime)
}
