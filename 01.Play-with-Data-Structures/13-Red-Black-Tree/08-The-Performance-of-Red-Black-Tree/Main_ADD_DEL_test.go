package RBTree

import (
	"fmt"
	"github.com/custergo/study_algo/01.Play-with-Data-Structures/13-Red-Black-Tree/08-The-Performance-of-Red-Black-Tree/BSTMap"
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestTreeAddDel(t *testing.T) {
	n := 20000000

	var testData []int
	for i := 0; i < n; i++ {
		testData = append(testData, rand.Intn(math.MaxInt64))
	}

	// Test BST
	startTime := time.Now()
	bst := BSTMap.NewBSTMap()
	for _, v := range testData {
		bst.Add(v, nil)
	}
	fmt.Println("BST: ", time.Now().Sub(startTime))

	// Test AVL
	startTime = time.Now()
	avl := BSTMap.NewBSTMap()
	for _, v := range testData {
		avl.Add(v, nil)
	}
	fmt.Println("AVL: ", time.Now().Sub(startTime))

	// Test RBTree
	startTime = time.Now()
	rbt := BSTMap.NewBSTMap()
	for _, v := range testData {
		rbt.Add(v, nil)
	}
	fmt.Println("BST: ", time.Now().Sub(startTime))
}

/**
BST:  1m33.704107917s
AVL:  1m47.615706207s
BST:  1m28.089209357s
*/
