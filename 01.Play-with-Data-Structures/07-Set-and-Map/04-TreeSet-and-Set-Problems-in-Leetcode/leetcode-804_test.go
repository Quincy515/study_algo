package _4_TreeSet_and_Set_Problems_in_Leetcode

import (
	"fmt"
	"testing"
)

func TestUniqueMorseRepresentations(t *testing.T) {
	words := []string{"gin", "zen", "gig", "msg"}
	fmt.Println(uniqueMorseRepresentations(words))
	fmt.Println(uniqueMorseRepresentations1(words))
	fmt.Println(uniqueMorseRepresentations2(words))
}
