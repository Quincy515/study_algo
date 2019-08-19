package HashTable

import (
	"fmt"
	"github.com/custergo/study_algo/1.Play-with-Data-Structures/14-Hash-Table/07-More-About-Resizing-in-Hash-Table/AVLTree"
	"github.com/custergo/study_algo/1.Play-with-Data-Structures/14-Hash-Table/07-More-About-Resizing-in-Hash-Table/BSTMap"
	"github.com/custergo/study_algo/1.Play-with-Data-Structures/14-Hash-Table/07-More-About-Resizing-in-Hash-Table/RBTree"
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
	fmt.Println("BST: ", diffTime) // BST:

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
	fmt.Println("AVL: ", diffTime) // AVL:

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
	fmt.Println("RBTree: ", diffTime) // RBTree:

	// Test HashTable
	startTime = time.Now()
	ht := NewHashTable() // Total words:  124448
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
