package Selection_Sort

import (
	"fmt"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	a := []int{10, 9, 9, 7, 6, 5, 4, 3, 2, 1}
	SelectionSort(a, len(a))
	fmt.Println(a)
}
