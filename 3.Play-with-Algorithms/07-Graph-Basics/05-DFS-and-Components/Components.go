package _5_DFS_and_Components

// 求无权图的联通分量
type Component struct {
	g       Graph  // 图的引用
	visited []bool // 记录dfs的过程中节点是否被访问
	ccount  int    // 记录联通分量个数
	id      []int  // 每个节点所对应的联通分量标记
}

// 构造函数,求出无权图的联通分量
func NewComponents(graph Graph) *Component {
	c := &Component{}
	c.g = graph
	c.visited = make([]bool, c.g.V())
	c.id = make([]int, c.g.V())
	c.ccount = 0
	for i := 0; i < c.g.V(); i++ {
		c.visited[i] = false
		c.id[i] = -1
	}

	// 求图的联通分量
	for i := 0; i < c.g.V(); i++ {
		if !c.visited[i] {
			c.dfs(i)
			c.ccount++
		}
	}
	return c
}

// 返回图的联通分量
func (this *Component) Count() int {
	return this.ccount
}

// 查询点v和点w是否联通
func (this *Component) IsConnected(v, w int) bool {
	assert(v >= 0 && v < this.g.V())
	assert(w >= 0 && w < this.g.V())
	return this.id[v] == this.id[w]
}

// 图的深度优先遍历
func (this *Component) dfs(v int) {
	this.visited[v] = true
	this.id[v] = this.ccount
	adj := NewAdjIteratorSparseGraph(this.g, v)

	for i := adj.Begin(); !adj.End(); i = adj.Next() {
		if !this.visited[i] {
			this.dfs(i)
		}
	}
}
