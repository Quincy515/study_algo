package _7_BFS_and_Shortest_Path

// 图的接口
type Graph interface {
	V() int
	E() int
	AddEdge(v, w int)
	Show()
}

func assert(exp bool) {
	if !exp { // 表示式为false
		panic("断言失败,发生错误!")
	}
}
