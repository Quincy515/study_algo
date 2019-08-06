package _1_Set_Basics_and_BSTSet

type Set interface {
	Add(e interface{})
	Remove(e interface{})
	Contains(e interface{}) bool
	GetSize() int
	IsEmpty() bool
}
