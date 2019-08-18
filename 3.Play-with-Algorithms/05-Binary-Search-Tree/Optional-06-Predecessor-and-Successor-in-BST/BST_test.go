package Binary_Search_Tree

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

// 测试二分搜索树中的predecessor和successor两个函数
func TestBST_ssor(t *testing.T) {
	assert := assert.New(t)
	// 生成0到n-1共n个数字的数组
	n := 1000
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i
	}
	// 将数组乱序
	rand.Shuffle(len(nums), func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})

	// 将这个n个数插入到二叉树中
	bst := NewBst()
	for i := 0; i < n; i++ {
		bst.Insert(i, i)
	}

	// 测试前驱算法,除了数字0没有前驱,每个数字x的前驱应该为x-1
	for i := 0; i < n; i++ {
		if i == 0 {
			assert.Equal(bst.predecessor(i), -9999)
			fmt.Println("The predecessor of 0 is NULL")
		} else {
			assert.Equal(bst.predecessor(i), i-1)
			fmt.Println("The predecessor of ", i, " is ", i-1)
		}
	}

	// 测试后继算法,除了数字没有n-1后继,每个数字x的后继应该为x+1
	for i := 0; i < n; i++ {
		if i == n-1 {
			assert.Equal(bst.successor(i), -9999)
			fmt.Println("The successor of ", i, "is NULL")
		} else {
			assert.Equal(bst.successor(i), i+1)
			fmt.Println("The successor of ", i, " is ", i+1)
		}
	}
}
