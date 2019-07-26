package _3_Single_Source_Path_BFS

type SingleSourcePath struct {
	G       *Graph
	s       int          // 源顶点
	pre     []int        //记录遍历的前一顶点
	visited map[int]bool // 是否被遍历过
}

func NewSingleSourcePath(G *Graph, s int) *SingleSourcePath {
	graphBfs := &SingleSourcePath{}
	graphBfs.G = G
	graphBfs.s = s
	graphBfs.visited = make(map[int]bool, G.V)
	graphBfs.pre = make([]int, G.V)
	for i := 0; i < G.V; i++ {
		graphBfs.pre[i] = -1
	}

	bfs(s, G, graphBfs.visited, graphBfs.pre) // 从源点s出发遍历该联通分量

	return graphBfs
}

func bfs(s int, self *Graph, visited map[int]bool, pre []int) {
	var queue []int          // 申请一个队列
	queue = append(queue, s) // 顶点s入队首
	visited[s] = true        // 入队标记为蓝色
	pre[s] = s               // 源点s的上一个顶点就记为s

	for len(queue) > 0 { // 只要队列不为空就继续执行循环
		v := queue[0]     // 取出队首元素顶点v
		queue = queue[1:] // 移除队首元素顶点

		for w := self.Adjacency(v).Front(); w != nil; w = w.Next() {
			ww := w.Value.(int) // 查看v顶点的所有相邻顶点
			if !visited[ww] {   // 查看顶点v的相邻顶点w是否已经被遍历过
				queue = append(queue, ww) // 没有被遍历过,就把w顶点入队
				visited[ww] = true        // 入队标记为蓝色,这样不会重复入队
				pre[ww] = v               // w顶点的上一个顶点是v
			}
		}
	}
}

func (s *SingleSourcePath) isConnectedTo(t int) bool {
	ValidateVertex(t, s.G.V)
	return s.visited[t] // 查看t是否被遍历到，表示和s在同一个联通分量，即有路径
}

func (s *SingleSourcePath) Path(target int) []int {
	var res []int
	if !s.isConnectedTo(target) {
		return res // 源s到目标t没有路径,返回空数组
	}

	cur := target // 目标顶点往回走到源s
	for cur != s.s {
		res = append(res, cur) // 把cur保存到路径中
		cur = s.pre[cur]       // cur赋值为前一个顶点
	}
	res = append(res, s.s) // 把源s添加到路径

	reverse(res) // 反转路径res
	return res
}

// 反转函数
func reverse(l []int) {
	for i := 0; i < int(len(l)/2); i++ {
		li := len(l) - i - 1
		l[i], l[li] = l[li], l[i]
	}
}
