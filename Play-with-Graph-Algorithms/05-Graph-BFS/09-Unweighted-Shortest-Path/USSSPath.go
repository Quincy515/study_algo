package _9_Unweighted_Shortest_Path

import "fmt"

// Unweighted Single Souce Shortest Path
// 无权图 单源 最短 路径
type USSSPath struct {
	G        *Graph
	s        int          // 源 顶点
	visited  map[int]bool // 是否被遍历过
	pre, dis []int
}

func NewUSSSPath(G *Graph, s int) *USSSPath {
	graphBfs := &USSSPath{}
	graphBfs.G = G
	graphBfs.s = s
	graphBfs.visited = make(map[int]bool, G.V)
	graphBfs.pre, graphBfs.dis = make([]int, G.V), make([]int, G.V)
	for i := 0; i < G.V; i++ {
		graphBfs.pre[i], graphBfs.dis[i] = -1, -1
	}
	bfs(s, G, graphBfs.visited, graphBfs.pre, graphBfs.dis)

	// 遍历所有的点，打印出来dis值
	for i := 0; i < G.V; i++ {
		fmt.Print(graphBfs.dis[i], " ") // 0 1 1 2 2 3 2
	}
	fmt.Println()
	return graphBfs
}

func bfs(s int, self *Graph, visited map[int]bool, pre, dis []int) {
	var queue []int          // 申请一个队列
	queue = append(queue, s) // 顶点s入队首
	visited[s] = true        // 入队标记为蓝色
	pre[s] = s               // 源点s的上一个顶点就记为s
	dis[s] = 0               // 源点s到顶点s距离为0

	for len(queue) > 0 { // 只要队列不为空就继续执行循环
		v := queue[0]     // 取出队首元素顶点v
		queue = queue[1:] // 移除队首元素顶点

		for w := self.Adjacency(v).Front(); w != nil; w = w.Next() {
			ww := w.Value.(int) // 查看v顶点的所有相邻顶点
			if !visited[ww] {   // 查看顶点v的相邻顶点w是否已经被遍历过
				queue = append(queue, ww) // 没有被遍历过,就把w顶点入队
				visited[ww] = true        // 入队标记为蓝色,这样不会重复入队
				pre[ww] = v               // w顶点的上一个顶点是v
				dis[ww] = dis[v] + 1      // 源点s到w的距离为源点s到w的上个v的距离+1
			}
		}
	}
}

func (u *USSSPath) IsConnectedTo(t int) bool {
	ValidateVertex(t, u.G.V)
	return u.visited[t]
}

func (u *USSSPath) Path(target int) []int {
	var res []int
	if !u.IsConnectedTo(target) {
		return res // 源s到目标t没有路径,返回空数组
	}

	cur := target // 目标顶点往回走到源s
	for cur != u.s {
		res = append(res, cur) // 把cur保存到路径中
		cur = u.pre[cur]       // cur赋值为前一个顶点
	}
	res = append(res, u.s) // 把源s添加到路径

	reverse(res) // 反转路径res
	return res
}

func (u *USSSPath) Dis(t int) int {
	ValidateVertex(t, u.G.V)
	return u.dis[t]
}

// 反转函数
func reverse(l []int) {
	for i := 0; i < int(len(l)/2); i++ {
		li := len(l) - i - 1
		l[i], l[li] = l[li], l[i]
	}
}
