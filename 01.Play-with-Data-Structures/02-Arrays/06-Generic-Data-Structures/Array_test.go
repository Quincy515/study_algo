package _6_Generic_Data_Structures

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	arr := NewArray(20)
	for i := 0; i < 10; i++ {
		arr.AddLast(i)
	}
	fmt.Println(arr)

	arr.Add(1, 100)
	fmt.Println(arr)

	arr.AddFirst(-1)
	fmt.Println(arr)

	arr.Remove(2)
	fmt.Println(arr)

	arr.RemoveElement(4)
	fmt.Println(arr)

	arr.RemoveFirst()
	fmt.Println(arr)

	// 测试student类型
	students := NewArray(10)

	students.AddLast(NewStudent("Alice", 100))
	students.AddLast(NewStudent("Bobb", 66))
	students.AddLast(NewStudent("Charlie", 88))
	fmt.Println(students)
}
