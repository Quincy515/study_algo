package AVLTree

import (
	"fmt"
	"github.com/custergo/study_algo/1.Play-with-Data-Structures/12-AVL-Tree/06-LR-and-RL/BSTMap"
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
	fmt.Println("BST: ", diffTime) // BST:  417.186376ms

	// Test AVL Tree
	startTime = time.Now()
	avl := NewAVLTree()
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
	fmt.Println("AVL: ", diffTime) // AVL:  344.215297ms
}
