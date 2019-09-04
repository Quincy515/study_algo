package LeetCode

import (
	"container/heap"
	"math/rand"
	"sort"
	"time"
)

// Sorting Time Complexity: O(nlogn) Space Complexity: O(1)
func findKthLargest1(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

// Using Min Heap to store the k largest elements
// Time Complexity: O(nlogk) Space Complexity: O(k)
// https://github.com/aQuaYi/LeetCode-in-Go/blob/master/Algorithms/0215.kth-largest-element-in-an-array/kth-largest-element-in-an-array.go
func findKthLargest2(nums []int, k int) int {
	temp := highHeap(nums)
	h := &temp
	heap.Init(h)
	if k == 1 {
		return (*h)[0]
	}
	for i := 1; i < k; i++ {
		heap.Remove(h, 0)
	}
	return (*h)[0]
}

// highHeap 实现了Heap 的接口
type highHeap []int

func (h highHeap) Len() int            { return len(h) }
func (h highHeap) Less(i, j int) bool  { return h[i] > h[j] } // h[i] > h[j] 使得 h[0] == max(h...)
func (h highHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *highHeap) Push(x interface{}) { *h = append(*h, x.(int)) } // Push 使用 *h，是因为增加了 h 的长度
func (h *highHeap) Pop() interface{} { // Pop 使用 *h，是因为减短了 h 的长度
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

// Based on Quick Sort Partition.Time Complexity: O(n) Space Complexity: O(logn)
func findKthLargest3(nums []int, k int) int {
	rand.Seed(time.Now().Unix())
	return _findKthLargest(nums, 0, len(nums)-1, k-1)
}

func _findKthLargest(nums []int, l, r, k int) int {
	if l == r {
		return nums[l]
	}
	p := partition(nums, l, r)
	if p == k {
		return nums[p]
	} else if k < p {
		return _findKthLargest(nums, l, p-1, k)
	} else { // k>p
		return _findKthLargest(nums, p+1, r, k)
	}
}

func partition(nums []int, l, r int) int {
	p := rand.Int()%(r-l+1) + l
	nums[l], nums[p] = nums[p], nums[l]

	lt := l + 1 // [l+1...lt) > p; [lt...i) < p
	for i := l + 1; i <= r; i++ {
		if nums[i] > nums[l] {
			nums[i], nums[lt] = nums[lt], nums[i]
			lt++
		}
	}
	nums[l], nums[lt-1] = nums[lt-1], nums[l]
	return lt - 1
}
