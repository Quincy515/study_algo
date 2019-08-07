package LeetCode_303

// sum[i]存储前i个元素和，sum[0] = 0
// 即sum[i]存储nums[0...i-1]的和
// sum(i, j) = sum[j + 1] - sum[i]
type NumArray struct {
	sum  []int
	data []int
}

func Constructor(nums []int) NumArray {
	sum := make([]int, len(nums)+1)
	sum[0] = 0

	for i := 1; i <= len(nums); i++ {
		sum[i] = sum[i-1] + nums[i-1]
	}

	return NumArray{sum: sum, data: nums}
}

func (this *NumArray) Update(i int, val int) {
	if i < 0 || i > len(this.data)-1 {
		panic("Index out of range.")
	}

	this.data[i] = val
	for i := i + 1; i <= len(this.data); i++ {
		this.sum[i] = this.data[i-1] + this.sum[i-1]
	}
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.sum[j+1] - this.sum[i]
}
