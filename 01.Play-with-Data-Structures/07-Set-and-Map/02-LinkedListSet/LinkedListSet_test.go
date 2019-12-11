package _2_LinkedListSet

import (
	"fmt"
	"testing"
)

func TestBSTSet(t *testing.T) {
	fmt.Println("pride-and-prejudice-傲慢与偏见")
	words1 := ReadFile("pride-and-prejudice.txt")
	fmt.Println("Total words: ", len(words1))

	set1 := NewLinkedListSet()
	for _, word := range words1 {
		set1.Add(word)
	}
	fmt.Println("Total different words: ", set1.GetSize())

	fmt.Println("a Tale of Two Cities-双城记")
	words2 := ReadFile("a-tale-of-two-cities.txt")
	fmt.Println("Total words: ", len(words2))
	set2 := NewLinkedListSet()
	for _, word := range words2 {
		set2.Add(word)
	}
	fmt.Println("Total different words: ", set2.GetSize())
}
