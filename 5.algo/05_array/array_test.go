package array

import (
	"testing"
)

func TestArray_Insert(t *testing.T) {
	capacity := 10
	arr := NewArray(uint(capacity))
	for i := 0; i < capacity-2; i++ {
		err := arr.Insert(uint(i), i+1)
		if nil != err {
			t.Fatal(err.Error())
		}
	}
	arr.Print() // |1|2|3|4|5|6|7|8

	arr.Insert(uint(6), 999)
	arr.Print() // |1|2|3|4|5|6|999|7|8

	arr.InsertToTail(666)
	arr.Print() // |1|2|3|4|5|6|999|7|8|666
}

func TestArray_Delete(t *testing.T) {
	capacity := 10
	arr := NewArray(uint(capacity))
	for i := 0; i < capacity; i++ {
		err := arr.Insert(uint(i), i+1)
		if nil != err {
			t.Fatal(err.Error())
		}
	}
	arr.Print()

	for i := 9; i >= 0; i-- {
		_, err := arr.Delete(uint(i))
		if nil != err {
			t.Fatal(err)
		}
		arr.Print()
	}
}

func TestArray_Find(t *testing.T) {
	capacity := 10
	arr := NewArray(uint(capacity))
	for i := 0; i < capacity; i++ {
		err := arr.Insert(uint(i), i+1)
		if nil != err {
			t.Fatal(err.Error())
		}
	}
	arr.Print()

	t.Log(arr.Find(0))  // 1 <nil>
	t.Log(arr.Find(9))  // 10 <nil>
	t.Log(arr.Find(11)) // 0 out of index range
}
