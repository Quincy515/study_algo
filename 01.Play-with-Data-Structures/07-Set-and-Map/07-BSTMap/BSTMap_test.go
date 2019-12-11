package _7_BSTMap

import (
	"fmt"
	"testing"
)

func TestBSTMap(t *testing.T) {
	fmt.Println("pride-and-prejudice")
	words := ReadFile("pride-and-prejudice.txt")
	fmt.Println("Total words: ", len(words))

	customMap := NewBSTMap()
	for _, word := range words {
		if customMap.Contains(word) {
			customMap.Set(word, customMap.Get(word).(int)+1)
		} else {
			customMap.Add(word, 1)
		}
	}

	fmt.Println("Total different words: ", customMap.GetSize())
	fmt.Println("Frequency of PRIDE: ", customMap.Get("pride"))
	fmt.Println("Frequency of PREJUDICE: ", customMap.Get("prejudice"))
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
