package Graph_Representation

// 图的接口
type Graph interface {
	V() int
	E() int
	AddEdge(v, w int)
}
