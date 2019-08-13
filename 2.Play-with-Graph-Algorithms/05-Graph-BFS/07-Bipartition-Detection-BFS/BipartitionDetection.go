package _7_Bipartition_Detection_BFS

type BipartitionDetection struct {
	G           *Graph
	visited     map[int]bool // 是否被遍历过
	colors      []int
	isBipartite bool
}

func NewBipartitionDetection(G *Graph) *BipartitionDetection {
	graphBfs := &BipartitionDetection{}
	graphBfs.G = G
	graphBfs.isBipartite = true
	graphBfs.visited = make(map[int]bool, G.V)
	graphBfs.colors = make([]int, G.V)
	for i := 0; i < G.V; i++ {
		graphBfs.colors[i] = -1
	}

	for v := 0; v < G.V; v++ {
		if !graphBfs.visited[v] {
			if !bfs(v, G, graphBfs.visited, graphBfs.colors) {
				graphBfs.isBipartite = false
				break
			}
		}
	}

	return graphBfs
}

func bfs(s int, self *Graph, visited map[int]bool, colors []int) bool {
	var queue []int          // 申请一个队列
	queue = append(queue, s) // 顶点s入队首
	visited[s] = true
	colors[s] = 0

	for len(queue) > 0 { // 只要队列不为空就继续执行循环
		v := queue[0]     // 取出队首元素顶点v
		queue = queue[1:] // 移除队首元素顶点

		// 查看v顶点的所有相邻顶点
		for w := self.Adjacency(v).Front(); w != nil; w = w.Next() {
			ww := w.Value.(int)
			if !visited[ww] { // 查看顶点v的相邻顶点w是否已经被遍历过
				queue = append(queue, ww) // 没有被遍历过,就把w顶点入队
				visited[ww] = true        // 入队标记为蓝色,这样不会重复入队
				colors[ww] = 1 - colors[v]
			} else if colors[v] == colors[ww] {
				return false
			}
		}
	}
	return true
}

func (b *BipartitionDetection) IsBipartite() bool {
	return b.isBipartite
}
