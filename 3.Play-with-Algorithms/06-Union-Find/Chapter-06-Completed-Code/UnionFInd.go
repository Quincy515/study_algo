package UnionFind

type UnionFind interface {
	IsConnected(int, int) bool
	UnionElements(int, int)
}
