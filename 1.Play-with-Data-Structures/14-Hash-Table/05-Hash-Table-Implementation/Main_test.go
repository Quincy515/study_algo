package HashTable

import (
	"fmt"
	"github.com/custergo/study_algo/Play-with-Data-Structures/13-Red-Black-Tree/08-The-Performance-of-Red-Black-Tree/AVLTree"
	"github.com/custergo/study_algo/Play-with-Data-Structures/13-Red-Black-Tree/08-The-Performance-of-Red-Black-Tree/BSTMap"
	"github.com/custergo/study_algo/Play-with-Data-Structures/14-Hash-Table/05-Hash-Table-Implementation/RBTree"
	"testing"
	"time"
)

func TestPProf(t *testing.T) {
	fmt.Println("pride-and-prejudice")
	words := ReadFile("pride-and-prejudice.txt")
	fmt.Println("Total words: ", len(words))

	// Test BST Tree
	startTime := time.Now()
	bst := BSTMap.NewBSTMap()
	for _, word := range words {
		if bst.Contains(word) {
			bst.Set(word, bst.Get(word).(int)+1)
		} else {
			bst.Add(word, 1)
		}
	}
	for _, word := range words {
		bst.Contains(word)
	}
	diffTime := time.Now().Sub(startTime)
	fmt.Println("BST: ", diffTime) // BST:  485.684543ms

	// Test AVL Tree
	startTime = time.Now()
	avl := AVLTree.NewAVLTree()
	for _, word := range words {
		if avl.Contains(word) {
			avl.Set(word, avl.Get(word).(int)+1)
		} else {
			avl.Add(word, 1)

		}
	}
	for _, word := range words {
		avl.Contains(word)
	}
	diffTime = time.Now().Sub(startTime)
	fmt.Println("AVL: ", diffTime) // AVL:  436.139659ms

	// Test RebBlack Tree
	startTime = time.Now()
	rbt := RBTree.NewRBTree()
	for _, word := range words {
		if rbt.Contains(word) {
			rbt.Set(word, rbt.Get(word).(int)+1)
		} else {
			rbt.Add(word, 1)

		}
	}
	for _, word := range words {
		rbt.Contains(word)
	}
	diffTime = time.Now().Sub(startTime)
	fmt.Println("RBTree: ", diffTime) // RBTree:  437.033303ms

	// Test HashTable
	startTime = time.Now()
	ht := NewHashTable(134448) // Total words:  124448
	for _, word := range words {
		if ht.Contains(word) {
			ht.Set(word, ht.Get(word).(int)+1)
		} else {
			ht.Add(word, 1)

		}
	}
	for _, word := range words {
		ht.Contains(word)
	}
	diffTime = time.Now().Sub(startTime)
	fmt.Println("HashTable: ", diffTime) //
}

/**
=== RUN   TestPProf
pride-and-prejudice
Total words:  124448
BST:  464.938819ms
AVL:  368.639172ms
RBTree:  380.410874ms
HashTable:  109.135234ms
--- PASS: TestPProf (1.38s)
PASS
*/
