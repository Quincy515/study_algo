package Bellman_Ford

import (
	"fmt"
	"strconv"
)

// Dijkstra算法求最短路径
type Dijkstra struct {
	g      Graph    // 图的引用
	s      int      // 单源最短路径的源，起始点
	distTo []weight // distTo[i]存储从起始点s到i的最短路径长度
	marked []bool   // 标记数组,在算法运行过程中标记节点i是否被访问
	from   []*Edge  // from[i]记录最短路径中,到达i点的边是哪一条
	// 可以用来恢复整个最短路径
}

// 构造函数,使用Dijkstra算法求最短路径
func NewDijkstra(graph Graph, s int) *Dijkstra {
	d := &Dijkstra{}
	// 算法初始化
	d.g = graph
	assert(s >= 0 && s < graph.V())
	d.s = s
	d.distTo = make([]weight, graph.V())
	d.marked = make([]bool, graph.V())
	d.from = make([]*Edge, graph.V())
	for i := 0; i < graph.V(); i++ {
		d.distTo[i] = 0.0
		d.marked[i] = false
		d.from[i] = nil
	}

	// Dijkstra
	// 使用索引堆记录当前找到的到达每个顶点的最短路径
	ipq := NewIndexMinHeap(graph.V())
	// 对于起始点s进行初始化
	d.distTo[s] = 0.0
	d.from[s] = NewEdge(s, s, 0.0)
	ipq.Insert(d.distTo[s], s)
	for !ipq.IsEmpty() {
		v := ipq.ExtractMinIndex()

		// distTo[v] 就是s到v的最短距离
		d.marked[v] = true

		// 对v的所有相邻节点进行松弛操作,进行更新
		adj := NewAdjIteratorSparseGraph(d.g, v) // 访问v节点的所有邻边
		for e := adj.Begin(); !adj.End(); e = adj.Next() {
			w := e.Other(v) // 邻边的另一端点w
			// 如果从s点到w点的最短路径还没有找到
			if !d.marked[w] { // 进行松弛判断
				temp, _ := strconv.ParseFloat(e.Wt().(string), 32)

				// 如果w点以前没有访问过
				// 或者访问过,但是通过当前的v点到w点距离更短,则进行更新
				if d.from[w] == nil || compare(add(d.distTo[v], temp), d.distTo[w].(float64)) < 0 {
					d.distTo[w] = add(d.distTo[v], temp)
					d.from[w] = e
					if ipq.contain(w) { // 如果最小索引堆中有w这个节点就更新
						ipq.change(w, d.distTo[w])
					} else { // 没有w节点就添加
						ipq.Insert(d.distTo[w], w)
					}
				}
			}
		}
	}
	return d
}

// 返回从s点到w点的最短路径长度
func (this Dijkstra) shortestPathTo(w int) weight {
	assert(w >= 0 && w < this.g.V())
	assert(this.hashPathTo(w))
	return this.distTo[w]
}

// 判断从s点到w点是否联通
func (this Dijkstra) hashPathTo(w int) bool {
	assert(w >= 0 && w < this.g.V())
	return this.marked[w]
}

// 寻找从s到w的最短路径,将整个路径经过的边存放在vec中
func (this Dijkstra) shortestPath(w int) []*Edge {
	assert(w >= 0 && w < this.g.V())
	assert(this.hashPathTo(w))

	// 通过from数组逆向查找到从s到w的路径,存放到栈中
	var stack []*Edge
	e := this.from[w]
	for e.V() != this.s {
		stack = append(stack, e)
		e = this.from[e.V()]
	}
	stack = append(stack, e)

	// 从栈中依次取出元素,获得顺序的从s到w的路径
	var res []*Edge
	for len(stack) > 0 {
		e = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, e)
	}
	return res
}

// 打印出从s点到w点的路径
func (this Dijkstra) showPath(w int) {
	assert(w >= 0 && w < this.g.V())
	assert(this.hashPathTo(w))

	path := this.shortestPath(w)
	for i := 0; i < len(path); i++ {
		fmt.Print(path[i].V(), " -> ")
		if i == len(path)-1 {
			fmt.Println(path[i].W())
		}
	}
}
