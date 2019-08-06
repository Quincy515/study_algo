package _8_More_about_Map

import (
	"fmt"
	_6_LinkedListMap "github.com/custergo/study_algo/Play-with-Data-Structures/07-Set-and-Map/06-LinkedListMap"
	_7_BSTMap "github.com/custergo/study_algo/Play-with-Data-Structures/07-Set-and-Map/07-BSTMap"
	"testing"
	"time"
)

func testMap(m Map, filename string) time.Duration {
	startTime := time.Now()

	words := ReadFile(filename)
	fmt.Println("Total words: ", len(words))

	for _, word := range words {
		if m.Contains(word) {
			m.Set(word, m.Get(word).(int)+1)
		} else {
			m.Add(word, 1)
		}
	}

	fmt.Println("Total different words: ", m.GetSize())
	fmt.Println("Frequency of PRIDE: ", m.Get("pride"))
	fmt.Println("Frequency of PREJUDICE: ", m.Get("prejudice"))
	return time.Now().Sub(startTime)
}

func TestMap(t *testing.T) {
	filename := "pride-and-prejudice.txt"

	bstMap := _7_BSTMap.NewBSTMap()
	time1 := testMap(bstMap, filename)
	fmt.Println("BST map: ", time1)

	linkedListMap := _6_LinkedListMap.NewLinkedListMap()
	time2 := testMap(linkedListMap, filename)
	fmt.Println("LinkedListMap map: ", time2)
}

/**
=== RUN   TestLinkedListMap
pride-and-prejudice
Total words:  124448
Total different words:  7018
Frequency of PRIDE:  50
Frequency of PREJUDICE:  11
--- PASS: TestLinkedListMap (12.72s)
PASS

=== RUN   TestBSTMap
pride-and-prejudice
Total words:  124448
Total different words:  7018
Frequency of PRIDE:  50
Frequency of PREJUDICE:  11
--- PASS: TestBSTMap (0.33s)
PASS
*/
