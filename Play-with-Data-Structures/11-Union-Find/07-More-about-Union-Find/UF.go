package UnionFind6

type UF interface {
	GetSize() int
	IsConnected(int, int) bool
	UnionElements(int, int)
}
