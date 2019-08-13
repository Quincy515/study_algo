package Selection_Sort

import (
	"fmt"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	a := []interface{}{10, 9, 9, 7, 6, 5, 4, 3, 2, 1}
	SelectionSort(a, len(a))
	fmt.Println(a) // [1 2 3 4 5 6 7 9 9 10]

	b := []interface{}{4.4, 3.3, 2.2, 1.1}
	SelectionSort(b, len(b))
	fmt.Println(b) // [1.1 2.2 3.3 4.4]

	c := []interface{}{"D", "C", "B", "A"}
	SelectionSort(c, len(c))
	fmt.Println(c) // [A B C D]

	d := Students{}
	d = append(d, Student{Name: "D", Score: 90})
	d = append(d, Student{Name: "C", Score: 100})
	d = append(d, Student{Name: "B", Score: 95})
	d = append(d, Student{Name: "A", Score: 95})

	StudentsSort(d)
	fmt.Println(d)
}
