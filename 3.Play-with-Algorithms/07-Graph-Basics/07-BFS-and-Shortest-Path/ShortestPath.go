package _7_BFS_and_Shortest_Path

import "fmt"

// 寻找无权图的最短路径
type ShortestPath struct {
	g       Graph  // 图的引用
	s       int    // 起始点
	visited []bool // 记录dfs的过程中节点是否被访问
	from    []int  // 记录路径,from[i]表示查找的路径上i的上一个节点
	ord     []int  // 记录路径中节点的次序.ord[i]表示i节点在路径中的次序
}

// 构造函数,寻找无权图graph从s点到其他点的最短路径
func NewShortestPath(graph Graph, s int) *ShortestPath {
	// 算法初始化
	assert(s >= 0 && s < graph.V())
	g := &ShortestPath{}
	g.g = graph
	g.s = s
	g.visited = make([]bool, graph.V())
	g.from = make([]int, graph.V())
	g.ord = make([]int, graph.V())
	for i := 0; i < graph.V(); i++ {
		g.visited[i] = false
		g.from[i] = -1
		g.ord[i] = -1
	}

	// 无向图最短路径算法,从s开始广度优先遍历整张图
	var queue []int
	queue = append(queue, s)
	g.visited[s] = true
	g.ord[s] = 0
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]

		adj := NewAdjIteratorSparseGraph(g.g, v)
		for i := adj.Begin(); !adj.End(); i = adj.Next() {
			if !g.visited[i] {
				queue = append(queue, i)
				g.visited[i] = true
				g.from[i] = v
				g.ord[i] = g.ord[v] + 1
			}
		}
	}

	return g
}

// 查询从s点到w点是否有路径
func (this *ShortestPath) hashPath(w int) bool {
	assert(w >= 0 && w < this.g.V())
	return this.visited[w]
}

// 查询从s点到w点的路径,存放在vec中
func (this *ShortestPath) path(w int, vec *[]int) {
	assert(w >= 0 && w < this.g.V())
	var stack []int
	// 通过from数组逆向查找到从s到w的路径,存放在栈中
	p := w
	for p != -1 {
		stack = append(stack, p)
		p = this.from[p]
	}

	// 从栈中依次取出元素,获得顺序的从s到w的路径
	// vec.clear()
	for len(stack) > 0 {
		*vec = append(*vec, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
}

// 打印出从s点到w点的路径
func (this *ShortestPath) showPath(w int) {
	assert(w >= 0 && w < this.g.V())
	var vec []int
	this.path(w, &vec)
	for i := 0; i < len(vec); i++ {
		fmt.Print(vec[i])
		if i == len(vec)-1 {
			fmt.Println()
		} else {
			fmt.Print(" -> ")
		}
	}
}

// 查看从s点到w点的最短路径长度
func (this *ShortestPath) length(w int) int {
	assert(w >= 0 && w < this.g.V())
	return this.ord[w]
}
