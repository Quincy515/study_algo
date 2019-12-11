package _2_BFS

type GraphBFS struct {
	G       *Graph
	visited map[int]bool // 是否被遍历过
	order   []int        // 存放广度优先遍历顺序
}

func NewGraphBFS(G *Graph) *GraphBFS {
	graphBfs := &GraphBFS{}
	graphBfs.G = G
	graphBfs.visited = make(map[int]bool, G.V)
	graphBfs.order = make([]int, 0)

	for v := 0; v < G.V; v++ {
		if !graphBfs.visited[v] {
			bfs(v, G, graphBfs.visited, &graphBfs.order)
		}
	}

	return graphBfs
}

func bfs(s int, self *Graph, visited map[int]bool, order *[]int) {
	var queue []int          // 申请一个队列
	queue = append(queue, s) // 顶点s入队首
	visited[s] = true        // 入队标记为蓝色

	for len(queue) > 0 { // 只要队列不为空就继续执行循环
		v := queue[0]              // 取出队首元素顶点v
		queue = queue[1:]          // 移除队首元素顶点
		*order = append(*order, v) // 添加到遍历结果中

		// 查看v顶点的所有相邻顶点
		for w := self.Adjacency(v).Front(); w != nil; w = w.Next() {
			ww := w.Value.(int)
			if !visited[ww] { // 查看顶点v的相邻顶点w是否已经被遍历过
				queue = append(queue, ww) // 没有被遍历过,就把w顶点入队
				visited[ww] = true        // 入队标记为蓝色,这样不会重复入队
			}
		}
	}
}

// 图的广度优先遍历结果
func (self *GraphBFS) Order() []int {
	return self.order
}
