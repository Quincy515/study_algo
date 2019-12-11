package Binary_Search_Tree

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 测试二分搜索树
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

	// order数组中存放[0...n)的所有元素
	order := make([]int, n)
	for i := 0; i < n; i++ {
		order[i] = i
	}

	// 打乱order数组的顺序
	rand.Shuffle(len(order), func(i, j int) {
		order[i], order[j] = order[j], order[i]
	})

	// 乱序删除[0...n)范围里的所有元素
	for i := 0; i < n; i++ {
		if bst.Contain(order[i]) {
			bst.Remove(order[i])
			fmt.Println("After remove ", order[i], " size= ", bst.Size())
		}
	}

	// 最终整个二分搜索树应该为空
	fmt.Println("最终整个二分搜索树应该为空", bst.Size())
}
