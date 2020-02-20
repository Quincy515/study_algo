package leetcode

// Based on Buckets
// each time, all we need to check is the bucket that x belongs to and its two adjacent buckets
//
// One thing worth mentioning is the difference from bucket sort –
// Each of our buckets contains at most one element at any time,
// because two elements in a bucket means "almost duplicate" and we can return early from the function.
// Therefore, a HashMap with an element associated with a bucket label is enough for our purpose.
//
// Time Complexity: O(n)
// Space Complexity: O(k)
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if t < 0 {
		return false
	}
	buckets := make(map[int]int)
	w := t + 1
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		id := getID(num, w)

		// check if bucket id is empty, each bucket may contain at most one element
		if _, ok := buckets[id]; ok {
			return true
		}

		// check the neighbor buckets for almost duplicate
		if _, ok := buckets[id-1]; ok && buckets[id-1] >= num-t {
			return true
		}

		if _, ok := buckets[id+1]; ok && buckets[id+1] <= num+t {
			return true
		}

		// now bucket id is empty and no almost duplicate in neighbor buckets
		buckets[id] = num

		if len(buckets) == k+1 {
			delete(buckets, getID(nums[i-k], w))
		}
	}
	return false
}

// Get the ID of the bucket from element value x and bucket width w
// Since `-3 / 5 = 0` and but we need `-3 / 5 = -1`.
func getID(x, w int) int {
	if x < 0 {
		return x + 1/w - 1
	}
	return x / w
}
