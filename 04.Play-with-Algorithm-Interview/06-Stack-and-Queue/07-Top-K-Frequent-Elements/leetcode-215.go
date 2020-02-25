package leetcode

import "container/heap"

// Time: O(n*log(k)), Space: O(k)
func findKthLargest(nums []int, k int) int {
	var h minHeap // 定义最小堆
	heap.Init(&h)
	for _, num := range nums { // 遍历数组
		if h.Len() < k { // 如果最小堆中的元素不足k个
			heap.Push(&h, num) // 向堆里添加元素即可
		} else if num > h[0] { // 否则说明堆中元素已经有k个,并且当新来的元素大于堆顶元素时
			heap.Pop(&h)       // 堆顶元素移除
			heap.Push(&h, num) // 加入新来的元素
			//heap.Fix(&h, num)
		}
	} // 数组遍历完后，堆中保存的就是top k元素
	return h[0] // 堆顶就是第k大的数字，直接返回即可
}

type minHeap []int // minHeap 实现了Heap 的接口

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i] < h[j] } // h[i] < h[j] 使得 h[0] == min(h...)
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(int)) } // Push 使用 *h，是因为增加了 h 的长度
func (h *minHeap) Pop() interface{} { // Pop 使用 *h，是因为减短了 h 的长度
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

// 辅助函数 输入数组和高低两个游标，输出是一次划分后分隔元素的下标，也就是pivot的下标
func partition(nums []int, low, high int) int {
	pivot, i, j := nums[low], low, high // 定义pivot和左右游标i,j,pivot这里简单的选取下标为low的元素
	for i < j {                         // 当i小于j时
		for i < j && nums[j] < pivot { // 如果nums[j]<pivot时
			j-- // 向左移动j
		}
		if i < j { // 循环结束后,如果i小于j,说明j指向的数字不小于pivot
			nums[i], nums[j] = nums[j], nums[i] // 交换i和j指向的数字
		}
		// 同理使用相同的方法处理游标i
		for i < j && nums[i] >= pivot { // 如果nums[i]>=pivot
			i++ // 不断向右移动i
		}
		if i < j { // 循环结束后,如果i小于j,说明i指向的数字小于pivot
			nums[i], nums[j] = nums[j], nums[i] // 交换i和j指向的数字
		}
	} // 注意这里每一步都加上了i<j的限定
	// 循环结束后i和j同时都指向中间的pivot
	return i // 返回i即可
}

// Time(avg): O(n), Time(worst): O(n^2), Space: O(1)
func findKthLargestK(nums []int, k int) int {
	low, high := 0, len(nums)-1 // 初始化高低两个游标
	for low <= high {
		// 调用partition函数把数组分成大小两个部分
		p := partition(nums, low, high) // 并返回分隔元素pivot的下标p
		if p == k-1 {                   // 如果p正好等于k-1
			return nums[p] // 说明下标p指向的元素既是第k大的元素
		} else if p > k-1 { //说明第k大的元素在p左边
			high = p - 1 // 更新高游标high
		} else { // p>k-1说明第k大的元素在p右边
			low = p + 1 // 更新低游标low
		}
	} // 循环结束后如果没有找到第k大的元素
	return -1 // 直接返回-1
}
