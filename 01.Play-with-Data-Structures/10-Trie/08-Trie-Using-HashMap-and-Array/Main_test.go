package Trie

import (
	"fmt"
	"github.com/custergo/study_algo/01.Play-with-Data-Structures/10-Trie/03-Searching-in-Trie/BSTSet"
	"testing"
	"time"
)

func TestTrie(t *testing.T) {
	fmt.Println("pride-and-prejudice-傲慢与偏见")
	words := ReadFile("pride-and-prejudice.txt")

	startTime := time.Now()

	bstSet := BSTSet.NewBSTSet()
	for _, word := range words {
		bstSet.Add(word)
	}
	for _, word := range words {
		bstSet.Contains(word)
	}
	diffTime := time.Now().Sub(startTime)
	fmt.Println("Total different words: ", bstSet.GetSize())
	fmt.Println("BSTSet: ", diffTime)

	// ---
	startTime = time.Now()
	trie := NewTrie()
	for _, word := range words {
		trie.Add(word)
	}
	for _, word := range words {
		trie.Contains(word)
	}
	diffTime = time.Now().Sub(startTime)
	fmt.Println("Total different words: ", trie.GetSize())
	fmt.Println("Trie: ", diffTime)

	// ---
	startTime = time.Now()
	trie3 := Constructor()
	for _, word := range words {
		trie3.Add(word)
	}
	for _, word := range words {
		trie3.Contains(word)
	}
	diffTime = time.Now().Sub(startTime)
	fmt.Println("Total different words: ", trie3.GetSize())
	fmt.Println("Array(Map) Trie: ", diffTime)
}

/**
=== RUN   TestTrie
pride-and-prejudice-傲慢与偏见
Total different words:  7018
BSTSet:  345.627818ms
Total different words:  7018
Trie:  83.736603ms
Total different words:  7018
Array(Map) Trie:  19.243324ms
--- PASS: TestTrie (0.50s)
PASS
*/
