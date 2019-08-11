package AVLMap

type Set interface {
	Add(e interface{})
	Remove(e interface{})
	Contains(e interface{}) bool
	GetSize() int
	IsEmpty() bool
}
