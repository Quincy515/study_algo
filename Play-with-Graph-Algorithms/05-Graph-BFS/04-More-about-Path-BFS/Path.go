package _4_More_about_Path_BFS

// 从s出发，当查询到用户要到达的t时，即使整张图还没有遍历完毕
// 也可以提前返回
import "fmt"

type Path struct {
	G       *Graph
	s, t    int    // 源点s和目标点t
	pre     []int  //记录遍历的前一顶点
	visited []bool // 是否被遍历过
}

func NewPath(G *Graph, s, t int) *Path {
	ValidateVertex(s, G.V) // 判断用户传入的顶点是否合法
	ValidateVertex(t, G.V)
	graphBfs := &Path{}
	graphBfs.G = G
	graphBfs.s = s
	graphBfs.t = t
	graphBfs.visited = make([]bool, G.V)
	for i := 0; i < G.V; i++ {
		graphBfs.visited[i] = false
	}
	graphBfs.pre = make([]int, G.V)
	for i := 0; i < G.V; i++ {
		graphBfs.pre[i] = -1
	}

	bfs(s, t, G, graphBfs.visited, graphBfs.pre)

	for _, e := range graphBfs.visited {
		fmt.Print(e, " ")
	}
	fmt.Println()

	return graphBfs
}

func bfs(s, t int, self *Graph, visited []bool, pre []int) {
	var queue []int          // 申请一个队列
	queue = append(queue, s) // 顶点s入队首
	visited[s] = true        // 入队标记为蓝色
	pre[s] = s               // 源点s的上一个顶点就记为s

	if s == t {
		return
	}

	for len(queue) > 0 { // 只要队列不为空就继续执行循环
		v := queue[0]     // 取出队首元素顶点v
		queue = queue[1:] // 移除队首元素顶点

		// 查看v顶点的所有相邻顶点
		for w := self.Adjacency(v).Front(); w != nil; w = w.Next() {
			ww := w.Value.(int)
			if !visited[ww] { // 查看顶点v的相邻顶点w是否已经被遍历过
				queue = append(queue, ww) // 没有被遍历过,就把w顶点入队
				visited[ww] = true        // 入队标记为蓝色,这样不会重复入队
				pre[ww] = v               // w顶点的上一个顶点是v
				if ww == t {
					return
				}
			}
		}
	}
}

func (p *Path) isConnectedTo(t int) bool {
	ValidateVertex(t, p.G.V)
	return p.visited[t] // 查看t是否被遍历到，表示和s在同一个联通分量，即有路径
}

func (p *Path) Path() []int {
	var res []int
	if !p.isConnectedTo(p.t) {
		return res // 源s到目标t没有路径,返回空数组
	}

	cur := p.t // 目标顶点往回走到源s
	for cur != p.s {
		res = append(res, cur) // 把cur保存到路径中
		cur = p.pre[cur]       // cur赋值为前一个顶点
	}
	res = append(res, p.s) // 把源s添加到路径

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
