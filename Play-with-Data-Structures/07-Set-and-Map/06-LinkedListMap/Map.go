package _6_LinkedListMap

type K interface{}
type V interface{}

type Map interface {
	Add(key K, value V)
	Remove(key K) V
	Contains(key K) bool
	Get(key K) V
	Set(key K, value V)
	GetSize() int
	IsEmpty() bool
}
