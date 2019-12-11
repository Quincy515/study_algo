package RBTree

import (
	"fmt"
	"github.com/custergo/study_algo/01.Play-with-Data-Structures/13-Red-Black-Tree/08-The-Performance-of-Red-Black-Tree/AVLTree"
	"github.com/custergo/study_algo/01.Play-with-Data-Structures/13-Red-Black-Tree/08-The-Performance-of-Red-Black-Tree/BSTMap"
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
	rbt := NewRBTree()
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
}
