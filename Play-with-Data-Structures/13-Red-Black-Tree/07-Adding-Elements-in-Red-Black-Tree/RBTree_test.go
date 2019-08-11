package RBTree

import (
	"fmt"
	"testing"
)

func TestAVLTree(t *testing.T) {
	fmt.Println("pride-and-prejudice")
	words := ReadFile("pride-and-prejudice.txt")
	fmt.Println("Total words: ", len(words))

	RBTree := NewRBTree()
	for _, word := range words {
		if RBTree.Contains(word) {
			RBTree.Set(word, RBTree.Get(word).(int)+1)
		} else {
			RBTree.Add(word, 1)
		}
	}

	fmt.Println("Total different words: ", RBTree.GetSize())
	fmt.Println("Frequency of PRIDE: ", RBTree.Get("pride"))
	fmt.Println("Frequency of PREJUDICE: ", RBTree.Get("prejudice"))
}
