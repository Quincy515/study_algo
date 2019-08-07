package LeetCode_303

type NumArray struct {
	segmentTree *SegmentTree
}

func Constructor(nums []int) NumArray {
	data := make([]interface{}, len(nums))
	for i := 0; i < len(nums); i++ {
		data[i] = nums[i]
	}

	segTree := NewSegmentTree(data, func(i, j interface{}) interface{} {
		return i.(int) + j.(int)
	})
	return NumArray{segTree}
}

func (this *NumArray) SumRange(i int, j int) int {
	if this.segmentTree == nil {
		panic("Segment Tree is null")
	}
	return this.segmentTree.Query(i, j).(int)
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
