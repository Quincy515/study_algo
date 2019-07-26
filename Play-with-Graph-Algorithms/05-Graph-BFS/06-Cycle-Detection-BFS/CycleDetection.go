package _6_Cycle_Detection_BFS

type CycleDetection struct {
	G         *Graph
	visited   map[int]bool // 是否被遍历过
	pre       []int
	hashCycle bool
}

func NewCycleDetection(G *Graph) *CycleDetection {
	graphBfs := &CycleDetection{}
	graphBfs.G = G
	graphBfs.visited = make(map[int]bool, G.V)
	graphBfs.pre = make([]int, G.V)
	for i := 0; i < G.V; i++ {
		graphBfs.pre[i] = -1
	}

	for v := 0; v < G.V; v++ {
		if !graphBfs.visited[v] {
			if bfs(v, G, graphBfs.visited, graphBfs.pre) {
				graphBfs.hashCycle = true
				break
			}
		}
	}

	return graphBfs
}

func bfs(s int, self *Graph, visited map[int]bool, pre []int) bool {
	var queue []int          // 申请一个队列
	queue = append(queue, s) // 顶点s入队首
	visited[s] = true        // 入队标记为蓝色
	pre[s] = s

	for len(queue) > 0 { // 只要队列不为空就继续执行循环
		v := queue[0]     // 取出队首元素顶点v
		queue = queue[1:] // 移除队首元素顶点

		// 查看v顶点的所有相邻顶点
		for w := self.Adjacency(v).Front(); w != nil; w = w.Next() {
			ww := w.Value.(int)
			if !visited[ww] { // 查看顶点v的相邻顶点w是否已经被遍历过
				queue = append(queue, ww) // 没有被遍历过,就把w顶点入队
				visited[ww] = true        // 入队标记为蓝色,这样不会重复入队
				pre[ww] = v
			} else if pre[v] != ww {
				return true
			}
		}
	}
	return false
}

func (c *CycleDetection) HasCycle() bool {
	return c.hashCycle
}
