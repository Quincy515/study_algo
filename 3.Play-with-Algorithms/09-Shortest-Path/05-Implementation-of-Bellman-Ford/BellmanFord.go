package Bellman_Ford

import (
	"fmt"
	"strconv"
)

// 使用BellmanFord算法求最短路径
type BellmanFord struct {
	g                 Graph    // 图的引用
	s                 int      // 起始点
	distTo            []weight // distTo[i]存储从起始点s到i的最短路径长度
	from              []*Edge  // from[i]记录最短路径中,到达i点的边是哪一条
	hashNegativeCycle bool     // 标记图中是否有负权环
}

// 构造函数,使用BellmanFord算法求最短路径
func NewBellmanFord(graph Graph, s int) *BellmanFord {
	b := &BellmanFord{}
	b.g = graph
	b.s = s
	b.distTo = make([]weight, graph.V())
	b.from = make([]*Edge, graph.V())
	// 初始化所有的节点s都不可达,有from数组来表示
	for i := 0; i < graph.V(); i++ {
		b.from[i] = nil
	}
	// 设置distTo[s]=0,并且让from[s]不为nil,表示初始s节点可达且距离为0
	b.distTo[s] = 0.0
	b.from[s] = NewEdge(s, s, 0.0)

	// Bellman-Ford的过程
	// 进行v-1次循环,每一次循环求出从起始点到其余所有点,最多使用pass步可到达的最短距离
	for pass := 1; pass < graph.V(); pass++ {
		// 每次循环中对所有的边进行一次松弛操作
		// 遍历所有边的方式是先遍历所有的顶点,然后遍历和所有顶点相邻的所有边
		for i := 0; i < graph.V(); i++ {
			// 使用实现的邻边迭代器遍历和所有顶点相邻的所有边
			adj := NewAdjIteratorSparseGraph(b.g, i)
			for e := adj.Begin(); !adj.End(); e = adj.Next() {
				temp, _ := strconv.ParseFloat(e.Wt().(string), 32)

				// 对于每一个边首先判断e->v()可达
				// 之后看如果e->w()以前没有到达过,显然可以更新distTo[e->w()]
				// 或者e->w()以前虽然到达过,但是通过这个e可以获得一个更短的距离,即可以进行一次松弛操作,可以更新distTo[e->w()]
				if b.from[e.V()] != nil && (b.from[e.W()] == nil ||
					compare(add(b.distTo[e.V()], temp), b.distTo[e.W()].(float64)) < 0) {
					b.distTo[e.W()] = add(b.distTo[e.V()], temp)
					b.from[e.W()] = e
				}
			}
		}
	}
	b.hashNegativeCycle = b.detectNegativeCycle()
	return b
}

// 判断图中是否有负权环
func (this BellmanFord) detectNegativeCycle() bool {
	for i := 0; i < this.g.V(); i++ {
		adj := NewAdjIteratorSparseGraph(this.g, i)
		for e := adj.Begin(); !adj.End(); e = adj.Next() {
			temp, _ := strconv.ParseFloat(e.Wt().(string), 32)
			if this.from[e.V()] != nil &&
				compare(add(this.distTo[e.V()], temp), this.distTo[e.W()].(float64)) < 0 {
				return true
			}
		}
	}
	return false
}

// 返回图中是否有负权环
func (this BellmanFord) negativeCycle() bool {
	return this.hashNegativeCycle
}

// 返回从s点到w点的最短路径长度
func (this BellmanFord) shortestPathTo(w int) weight {
	assert(w >= 0 && w < this.g.V())
	assert(!this.hashNegativeCycle)
	assert(this.hasPathTo(w))
	return this.distTo[w]
}

// 判断从s点到w点是否联通
func (this BellmanFord) hasPathTo(w int) bool {
	assert(w >= 0 && w < this.g.V())
	return this.from[w] != nil
}

// 寻找从s点到w点的最短路径,将整个路径经过的边存放到vec中
func (this BellmanFord) shortestPath(w int) []*Edge {
	assert(w >= 0 && w < this.g.V())
	assert(!this.hashNegativeCycle)
	assert(this.hasPathTo(w))

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

// 打印出从s点w点的路径
func (this BellmanFord) showPath(w int) {
	assert(w >= 0 && w < this.g.V())
	assert(!this.hashNegativeCycle)
	assert(this.hasPathTo(w))

	res := this.shortestPath(w)
	for i := 0; i < len(res); i++ {
		fmt.Print(res[i].V(), " -> ")
		if i == len(res)-1 {
			fmt.Println(res[i].W())
		}
	}
}
