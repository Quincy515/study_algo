package _5_CC_BFS

// 使用BFS解决联通分量问题
type CC struct {
	G       *Graph
	visited []int // 是否被遍历过
	cccount int
}

func NewCC(G *Graph) *CC {
	graphBfs := &CC{}
	graphBfs.G = G
	graphBfs.visited = make([]int, G.V)
	for i := 0; i < G.V; i++ {
		graphBfs.visited[i] = -1
	}

	for v := 0; v < G.V; v++ {
		if graphBfs.visited[v] == -1 {
			bfs(v, graphBfs.cccount, G, graphBfs.visited)
			graphBfs.cccount++
		}
	}

	return graphBfs
}

func bfs(s, ccid int, self *Graph, visited []int) {
	var queue []int          // 申请一个队列
	queue = append(queue, s) // 顶点s入队首
	visited[s] = ccid        // 入队标记为蓝色

	for len(queue) > 0 { // 只要队列不为空就继续执行循环
		v := queue[0]     // 取出队首元素顶点v
		queue = queue[1:] // 移除队首元素顶点

		// 查看v顶点的所有相邻顶点
		for w := self.Adjacency(v).Front(); w != nil; w = w.Next() {
			ww := w.Value.(int)
			if visited[ww] == -1 { // 查看顶点v的相邻顶点w是否已经被遍历过
				queue = append(queue, ww) // 没有被遍历过,就把w顶点入队
				visited[ww] = ccid        // 入队标记为蓝色,这样不会重复入队
			}
		}
	}
}

func (c *CC) Count() int {
	return c.cccount
}

func (c *CC) isConnected(v, w int) bool {
	ValidateVertex(v, c.G.V)
	ValidateVertex(w, c.G.V)
	return c.visited[v] == c.visited[w]
}

func (c *CC) components() [][]int {
	var res [][]int
	for i := 0; i < c.cccount; i++ {
		tmp := make([]int, 0)
		res = append(res, tmp)
	}

	for v := 0; v < c.G.V; v++ {
		res[c.visited[v]] = append(res[c.visited[v]], v)
	}

	return res
}
