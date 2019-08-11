package RBTree

import (
	"fmt"
	"github.com/custergo/study_algo/Play-with-Data-Structures/13-Red-Black-Tree/08-The-Performance-of-Red-Black-Tree/BSTMap"
	"testing"
	"time"
)

func TestTree(t *testing.T) {
	n := 20000000

	var testData []int
	for i := 0; i < n; i++ {
		testData = append(testData, i)
	}

	// Test AVL
	startTime := time.Now()
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
AVL:  1m47.615706207s
BST:  1m28.089209357s
*/
