package LeetCode_380

import (
	"math/rand"
	"strconv"
	"time"
)

type RandomizedSet struct {
	m    map[string]int
	nums []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	rand.Seed(time.Now().Unix())
	return RandomizedSet{
		m: make(map[string]int),
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	key := strconv.Itoa(val)
	if _, ok := this.m[key]; ok {
		return false
	}
	this.nums = append(this.nums, val)
	index := len(this.nums) - 1
	this.m[key] = index
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	key := strconv.Itoa(val)
	if _, ok := this.m[key]; !ok {
		return false
	}
	index := this.m[key]
	delete(this.m, key)

	num := this.nums[len(this.nums)-1]
	this.nums = this.nums[:len(this.nums)-1]

	if num != val {
		this.nums[index] = num
		this.m[strconv.Itoa(num)] = index
	}
	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	index := rand.Intn(len(this.nums))
	return this.nums[index]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
