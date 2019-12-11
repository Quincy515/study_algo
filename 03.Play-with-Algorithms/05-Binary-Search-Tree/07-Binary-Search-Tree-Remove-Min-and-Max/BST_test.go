package Binary_Search_Tree

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 测试二分搜索树中的removeMin和removeMax
func TestBST_Remove(t *testing.T) {
	rand.Seed(time.Now().Unix())
	bst := NewBst()

	// 取n个取值范围在[0...m)的随机整数放进二分搜索树中
	n, m := 100, 100
	for i := 0; i < n; i++ {
		key := rand.Int() % m
		// 为了后序测试方便,这里value值取和key值一样
		value := key
		fmt.Print(key, " ")
		bst.Insert(key, value)
	}
	fmt.Println()
	// 注意,由于随机生成的数据有重复,所以bst中的数据数量大概率是小于n的

	// 测试 removeMin
	// 输出的元素应该是从小到大排列的
	fmt.Println("Test removeMin: ")
	for !bst.IsEmpty() {
		fmt.Print("min: ", bst.Minimum(), " , ")
		bst.RemoveMin()
		fmt.Println("After removeMin, size= ", bst.Size())
	}

	for i := 0; i < n; i++ {
		key := rand.Int() % n
		// 为了后序测试方便,这里value值取的和key值一样
		value := key
		bst.Insert(key, value)
	}
	// 注意,由于随机生成的数据有重复,所以bst中的数据数量大概率是小于n的

	// 测试 removeMax
	// 输出的元素应该是从大到小排列的
	fmt.Println("Test removeMax: ")
	for !bst.IsEmpty() {
		fmt.Print("max: ", bst.Maximum(), " , ")
		bst.RemoveMax()
		fmt.Println("After removeMax, size= ", bst.Size())
	}

}
