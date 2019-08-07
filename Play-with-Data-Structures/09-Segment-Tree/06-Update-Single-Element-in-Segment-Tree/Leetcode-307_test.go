package _6_Update_Single_Element_in_Segment_Tree

import (
	"fmt"
	"testing"
)

func TestSumRange(t *testing.T) {
	nums := []int{1, 3, 5}
	obj := Constructor(nums)

	fmt.Println(obj.SumRange(0, 2)) // 9
	obj.Update(1, 2)
	fmt.Println(obj.SumRange(0, 2)) // 8
}
