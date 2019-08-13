package _1_Remove_Min_and_Max_in_BST

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestBST(t *testing.T) {
	bst := NewBST()

	n := 1000

	// test removeMin
	var nums []interface{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		bst.Add(rand.Intn(10000))
	}

	for !bst.IsEmpty() {
		nums = append(nums, bst.RemoveMin())
	}
	fmt.Println(nums)

	for i := 1; i < len(nums); i++ {
		if compare(nums[i-1], nums[i]) > 0 {
			panic("Error")
		}
	}
	fmt.Println("removeMin test completed.")

	// test removeMax
	for i := 0; i < n; i++ {
		bst.Add(rand.Intn(10000))
	}
	var numMax []interface{}
	for !bst.IsEmpty() {
		numMax = append(numMax, bst.RemoveMax())
	}
	fmt.Println(numMax)

	for i := 1; i < len(numMax); i++ {
		if compare(numMax[i-1], numMax[i]) < 0 {
			panic("Error")
		}
	}
	fmt.Println("removeMax test completed.")
}
