package Binary_Search_Tree

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestBST_Traverse(t *testing.T) {
	// 测试二分搜索树的前中后序遍历
	rand.Seed(time.Now().Unix())
	bst := NewBst()

	// 取n个取值范围在[0...m)的随机整数放进二分搜索树中
	n, m := 10, 100
	for i := 0; i < n; i++ {
		key := rand.Int() % m
		// 为了后序测试方便,这里value值取和key值一样
		value := key
		fmt.Print(key, " ")
		bst.Insert(key, value)
	}
	fmt.Println()

	// 测试二分搜索树的size
	fmt.Println("size: ", bst.Size())

	// 测试二分搜索树的前序遍历 preOrder
	fmt.Print("preOrder: ")
	bst.PreOrder()
	fmt.Println()

	// 测试二分搜索树的中序遍历 inOrder
	fmt.Println("inOrder: ")
	bst.InOrder()
	fmt.Println()

	// 测试二分搜索树的后序遍历 postOrder
	fmt.Println("postOrder: ")
	bst.PostOrder()
	fmt.Println()
}
