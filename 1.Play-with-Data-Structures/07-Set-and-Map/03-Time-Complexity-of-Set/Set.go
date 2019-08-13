package _3_Time_Complexity_of_Set

type Set interface {
	Add(e interface{})
	Remove(e interface{})
	Contains(e interface{}) bool
	GetSize() int
	IsEmpty() bool
}
