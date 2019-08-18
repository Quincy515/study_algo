package Floor_and_Ceil_in_Binary_Search

import (
	"fmt"
	"testing"
)

// 测试用二分查找法实现的floor和ceil两个函数
// 仔细观察在我们的测试用例中,有若干的重复元素,对于这些重复元素,floor和ceil计算结果的差别
func TestBinarySearchFloorCeil(t *testing.T) {
	a := []int{1, 1, 1, 2, 2, 2, 2, 2, 4, 4, 5, 5, 5, 6, 6, 6}
	n := len(a)
	for i := 0; i <= 8; i++ {
		floorIndex := floor(a, n, i)
		fmt.Println("the floor index of ", i, " is ", floorIndex)

		if floorIndex >= 0 && floorIndex < n {
			fmt.Println("The value is ", a[floorIndex], ".")
		}

		ceilIndex := ceil(a, len(a), i)
		fmt.Println("the ceil index of ", i, " is ", ceilIndex, ".")
		if ceilIndex >= 0 && ceilIndex < n {
			fmt.Println("The value is ", a[ceilIndex], ".")
		}
	}
}
