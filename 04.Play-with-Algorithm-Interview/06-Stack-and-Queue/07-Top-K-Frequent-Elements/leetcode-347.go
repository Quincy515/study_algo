package leetcode

import "container/heap"

// Time: O(n*log(k)), Space: O(n)
func topKFrequent1(nums []int, k int) []int {
	freqMap := make(map[int]int, 0) // 定义一个哈希表记录每个数字以及它的出现频率
	for _, num := range nums {      // 遍历数组中的每个数字num
		freqMap[num]++
	}
	var pq minHeap // 定义一个最小堆
	heap.Init(&pq)
	for key, val := range freqMap { // 遍历哈希表
		heap.Push(&pq, &Node{key, val}) // 把哈希表中每个元素加入到最小堆中
		if pq.Len() > k {               // 如果最小堆的大小已经大于k
			heap.Pop(&pq) // 移除堆顶元素,即移除当前堆中出现频率最小的元素
		}
	}
	var result []int   // 定义结果列表
	for pq.Len() > 0 { // 把最小堆中的数字加入到结果列表
		tmp := heap.Pop(&pq).(*Node)
		result = append(result, tmp.Val)
	}
	return result // 最后返回结果列表即可
}

type Node struct {
	Val, Freq int
}

type minHeap []*Node

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i].Freq < h[j].Freq }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(*Node)) }
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 辅助函数，输入元组列表和高低两个游标，输出是一次划分后把数组分隔是大小频率两部分的分隔元素下标，也就是基准元素pivot下标
func partition(entries map[int]int, low, high int) int {
	pivot := entries[low] // 首先把下标为low的元素的出现频率取出作为基准值
	i, j := low, high     // 初始化i为low，j为high
	for i < j {           // 当i小于j时不断执行以下操作
		for i < j && entries[j] < pivot { // 如果j指向的元素出现的频率小于基准值
			j-- // 向左移动j
		} // 循环结束后如果i仍然小于j，则说明此时j指向的元素的出现频率不小于基准值
		if i < j { // 于是交换i和j指向的元素
			entries[i], entries[j] = entries[j], entries[i]
		}
		// 同理使用相同的方法处理游标i
		for i < j && entries[i] >= pivot { // 如果i指向的元素出现的频率大于基准值
			i++ // 向右移动i
		} // 循环结束后如果i仍然小于j，则说明此时i指向的元素的出现频率小于基准值
		if i < j { // 于是交换i和j指向的元素
			entries[i], entries[j] = entries[j], entries[i]
		}
	} // 注意这里每一步都加上了i<j的限定
	// 循环结束后i和j同时都指向中间的pivot
	return i // 返回i即可
}

// Time(avg): O(n), Time(worst): O(n^2), Space: O(n)
func topKFrequent2(nums []int, k int) []int {
	freqMap := make(map[int]int, 0) // 定义一个哈希表记录每个数字以及它的出现频率
	for _, num := range nums {      // 遍历数组
		freqMap[num]++ // 统计每个数字的出现频率
	}
	entries:=// 哈希表转成元组列表
	low, high := 0, len(freqMap)-1 // 初始化高低两个游标low,high
	for low < high {
		p := partition(entries, low, high)
		if p == k-1 {                   // 如果p正好等于k-1
			break // 说明下标p指向的元素既是第k大的元素
		} else if p > k-1 { //说明第k大的元素在p左边
			high = p - 1 // 更新高游标high
		} else { // p>k-1说明第k大的元素在p右边
			low = p + 1 // 更新低游标low
		}
	} // 循环结束后如果没有找到第k大的元素
	var result []int
	for i:=0;i<k;i++{
		result = append(result, entries[i].getKey())
	}
	return result
}

// Time: O(n), Space: O(n)
func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int, 0) // 定义一个哈希表记录每个数字以及它的出现频率
	for _, num := range nums {      // 遍历数组
		freqMap[num]++ // 统计每个数字的出现频率
	}

	buckets := make([][]int, len(nums)+1) // 定义一个大小为n+1的列表，表示n+1个桶，每个桶也是一个列表，记录出现频率相同的数字
	for i := 0; i <= len(nums); i++ {     // n+1个桶初始化为空的列表
		buckets[i] = make([]int, 0)
	}

	for k, v := range freqMap { // 遍历哈希表中的元素
		buckets[v] = append(buckets[v], k) // 根据出现频率取出相应的桶,然后把数字加入桶中
	}
	var result []int // 定义结果列表
	// 从后向前遍历每个桶，同时判断k大于0,才继续执行
	for i := len(buckets) - 1; i >= 0 && k > 0; i-- {
		bucket := buckets[i]                        // 把第i个桶取出
		for j := 0; j < len(bucket) && k > 0; j++ { // 遍历这个桶内的数字,并且k>0,才继续执行
			result = append(result, bucket[j]) // 把桶内的数字加入到结果列表
			k--                                // 并且k-1
		}
	}
	return result // 循环结束后返回结果列表即可
}

