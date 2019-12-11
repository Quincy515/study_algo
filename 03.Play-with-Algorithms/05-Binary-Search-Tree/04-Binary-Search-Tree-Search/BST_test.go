package Binary_Search_Tree

import (
	"math/rand"
	"testing"
)

func TestBST(t *testing.T) {
	n := 1000000

	// 创建一个数组,包含[0...n)的所有元素
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}

	// 打乱数组顺序
	for i := 0; i < n; i++ {
		pos := rand.Int() % n // [0,n)范围
		arr[i], arr[pos] = arr[pos], arr[i]
	}

	// 由于实现的二分搜索树不是平衡二叉树
	// 所以如果按照顺序插入一组数据,二分搜索树会退化成为一个链表
	// 平衡二叉树的实现,在这里没有涉及

	bst := NewBst()
	for i := 0; i < n; i++ {
		bst.Insert(arr[i], arr[i])
	}

	// 对[0...2*n)的所有整型测试在二分搜索树中查找
	// 若i在[0...n)之间,则能查找到整型所对应的字符串
	// 若i在[n...2*n)之间,则结果为nil
	for i := 0; i < 2*n; i++ {
		res := bst.Search(i)
		if i < n {
			if res != i {
				panic("二分搜索树失败!")
			}
		} else {
			if res != 0 {
				panic("二分搜索树失败!!")
			}
		}
	}
}
