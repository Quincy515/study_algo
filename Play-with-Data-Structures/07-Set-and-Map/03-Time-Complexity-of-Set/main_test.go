package _3_Time_Complexity_of_Set

import (
	"fmt"
	_1_Set_Basics_and_BSTSet "github.com/custergo/study_algo/Play-with-Data-Structures/07-Set-and-Map/01-Set-Basics-and-BSTSet"
	_2_LinkedListSet "github.com/custergo/study_algo/Play-with-Data-Structures/07-Set-and-Map/02-LinkedListSet"
	"testing"
	"time"
)

func testSet(set Set, filename string) time.Duration {
	startTime := time.Now()

	words := ReadFile(filename)
	fmt.Println("Total words: ", len(words))
	for _, word := range words {
		set.Add(word)
	}
	fmt.Println("Total different words: ", set.GetSize())
	return time.Now().Sub(startTime)
}

func TestSetTime(t *testing.T) {
	filename := "pride-and-prejudice.txt"
	bstSet := _1_Set_Basics_and_BSTSet.NewBSTSet()
	time1 := testSet(bstSet, filename)
	fmt.Println("BST set: ", time1)

	linkedListSet := _2_LinkedListSet.NewLinkedListSet()
	time2 := testSet(linkedListSet, filename)
	fmt.Println("linkedList set: ", time2)
}
