package AVLMap

import (
	"fmt"
	"github.com/custergo/study_algo/Play-with-Data-Structures/12-AVL-Tree/08-Map-and-Set-in-AVL-Tree/BSTMap"
	"github.com/custergo/study_algo/Play-with-Data-Structures/12-AVL-Tree/08-Map-and-Set-in-AVL-Tree/BSTSet"
	"github.com/custergo/study_algo/Play-with-Data-Structures/12-AVL-Tree/08-Map-and-Set-in-AVL-Tree/LinkedListMap"
	"github.com/custergo/study_algo/Play-with-Data-Structures/12-AVL-Tree/08-Map-and-Set-in-AVL-Tree/LinkedListSet"
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

	bstMap := BSTMap.NewBSTMap()
	time1 := testMap(bstMap, filename)
	fmt.Println("BST map: ", time1)
	fmt.Println()

	linkedListMap := LinkedListMap.NewLinkedListMap()
	time2 := testMap(linkedListMap, filename)
	fmt.Println("LinkedListMap map: ", time2)
	fmt.Println()

	avlMap := NewAVLMap()
	time3 := testMap(avlMap, filename)
	fmt.Println("AVL Map: ", time3)

}

/**
=== RUN   TestMap
Total words:  124448
Total different words:  7018
Frequency of PRIDE:  50
Frequency of PREJUDICE:  11
BST map:  362.929152ms

Total words:  124448
Total different words:  7018
Frequency of PRIDE:  50
Frequency of PREJUDICE:  11
LinkedListMap map:  15.449438764s

Total words:  124448
Total different words:  7018
Frequency of PRIDE:  50
Frequency of PREJUDICE:  11
AVL Map:  315.644255ms
--- PASS: TestMap (16.13s)
PASS
*/

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
	bstSet := BSTSet.NewBSTSet()
	time1 := testSet(bstSet, filename)
	fmt.Println("BST set: ", time1)

	linkedListSet := LinkedListSet.NewLinkedListSet()
	time2 := testSet(linkedListSet, filename)
	fmt.Println("linkedList set: ", time2)

	avlSet := NewAVLSet()
	time3 := testSet(avlSet, filename)
	fmt.Println("AVL Set: ", time3)
}

/**
=== RUN   TestSetTime
Total words:  124448
Total different words:  7018
BST set:  155.09073ms
Total words:  124448
Total different words:  7018
linkedList set:  4.427574336s
Total words:  124448
Total different words:  7018
AVL Set:  126.833192ms
--- PASS: TestSetTime (4.71s)
PASS
*/
