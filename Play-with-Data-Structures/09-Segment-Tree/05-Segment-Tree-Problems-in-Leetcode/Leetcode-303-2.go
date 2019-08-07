package LeetCode_303

// sum[i]存储前i个元素和，sum[0] = 0
// 即sum[i]存储nums[0...i-1]的和
// sum(i, j) = sum[j + 1] - sum[i]
type NumArray struct {
	sum []int
}

func Constructor(nums []int) NumArray {
	sum := make([]int, len(nums)+1)
	sum[0] = 0

	for i := 1; i <= len(nums); i++ {
		sum[i] = sum[i-1] + nums[i-1]
	}

	return NumArray{sum}
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.sum[j+1] - this.sum[i]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
