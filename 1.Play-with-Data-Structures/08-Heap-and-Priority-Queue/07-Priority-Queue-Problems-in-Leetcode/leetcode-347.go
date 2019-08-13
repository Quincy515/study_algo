package Leetcode_347

func TopKFrequent(nums []int, k int) []int {
	m := make(map[int]int) // 保存每个元素和它的频次
	for _, num := range nums {
		m[num]++
	}
	// 求出前K个频次最高的元素
	pq := NewPriorityQueue()
	for num, freq := range m {
		if pq.GetSize() < k {
			pq.Enqueue(&Freq{num, freq})
		} else if freq > pq.GerFront().(*Freq).freq {
			pq.Dequeue()
			pq.Enqueue(&Freq{num, freq})
		}
	}

	var res []int
	for !pq.IsEmpty() {
		res = append(res, pq.Dequeue().(*Freq).e)
	}
	return res
}
