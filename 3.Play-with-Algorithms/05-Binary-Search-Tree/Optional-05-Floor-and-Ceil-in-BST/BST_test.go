package Binary_Search_Tree

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

// 测试二分搜索树中的floor和ceil两个函数
func TestBST_Remove(t *testing.T) {
	assert := assert.New(t)
	bst := NewBst()

	// 将[0,n)之间的偶数保存在nums中
	n := 1000
	var nums []int
	for i := 0; i < n; i += 2 {
		nums = append(nums, i)
	}
	minNum := nums[0]
	maxNum := nums[len(nums)-1]

	// 将nums乱序处理
	rand.Shuffle(len(nums), func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})

	// 向二分搜索树中插入[0,n)之间的所有偶数
	for _, num := range nums {
		bst.Insert(num, num)
	}

	// 对[0...n]区间里的n+1个数,调用二分搜索树的floor和ceil,查看其结果
	for i := 0; i < n; i++ {
		// 测试floor
		floorKey := bst.Floor(i)
		if i%2 == 0 {
			if i >= 0 && i < n {
				assert.Equal(floorKey, i)
			} else if i < 0 {
				assert.Equal(floorKey, -9999)
			} else {
				assert.Equal(floorKey, maxNum)
			}
		} else {
			if i-1 >= 0 && i-1 < n {
				assert.Equal(floorKey, i-1)
			} else if i-1 < 0 {
				assert.Equal(floorKey, -9999)
			} else {
				assert.Equal(floorKey, maxNum)
			}
		}

		fmt.Print("The floor of ", i, " is ")
		if floorKey == -9999 {
			fmt.Println("NULL")
		} else {
			fmt.Println(floorKey)
		}

		// 测试ceil
		ceilKey := bst.Ceil(i)
		if i%2 == 0 {
			if i >= 0 && i < n {
				assert.Equal(ceilKey, i)
			} else if i < 0 {
				assert.Equal(ceilKey, minNum)
			} else {
				assert.Equal(ceilKey, -9999)
			}
		} else {
			if i+1 >= 0 && i+1 < n {
				assert.Equal(ceilKey, i+1)
			} else if i+1 < 0 {
				assert.Equal(ceilKey, minNum)
			} else {
				assert.Equal(ceilKey, -9999)
			}
		}
		fmt.Print("the ceil of ", i, " is ")
		if ceilKey == -9999 {
			fmt.Println("NULL")
		} else {
			fmt.Println(ceilKey)
		}
	}
}
