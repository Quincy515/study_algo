package _5_DFS_and_Components

import "fmt"

// 稀疏图-邻接表
type SparseGraph struct {
	n        int     // 节点数
	m        int     // 边数
	directed bool    // 是否为有向图
	g        [][]int // 图的具体数据
}

func NewSparseGraph(n int, directed bool) *SparseGraph {
	// g初始化为n个空的vector,表示每一个g[i]都为空,即没有任何边
	// m: 0 表示初始化没有任何边
	return &SparseGraph{n: n, m: 0, directed: directed, g: make([][]int, n)}
}

// 返回节点个数
func (this *SparseGraph) V() int {
	return this.n
}

// 返回边的个数
func (this *SparseGraph) E() int {
	return this.m
}

// 向图中添加一个边
func (this *SparseGraph) AddEdge(v, w int) {
	if v < 0 || v >= this.n || w < 0 || w >= this.n {
		panic("向图中添加一条边超出了图的最大范围")
	}
	this.g[v] = append(this.g[v], w) // 在v中添加w,表示v和w相连
	if v != w && !this.directed {    // 不是自环边并且是无向图
		this.g[w] = append(this.g[w], v) // 在w中添加v,表示w和v相连
	}
	this.m++ // 添加边的数目
	// 不去考虑平行边,如果考虑平行边,需要判断v和w之间是否有边,时间复杂度就变成了O(n)
}

// 验证图中是否有从v到w的边-时间复杂度在最差情况下是O(n)
func (this SparseGraph) hasEdge(v, w int) bool {
	if v < 0 || v >= this.n || w < 0 || w >= this.n {
		panic("向图中添加一条边超出了图的最大范围")
	}
	for i := 0; i < len(this.g[v]); i++ {
		if this.g[v][i] == w {
			return true
		}
	}
	return false
}

// 显示图的信息
func (this *SparseGraph) Show() {
	for i := 0; i < this.n; i++ {
		fmt.Print("vertex ", i, ":\t")
		for j := 0; j < len(this.g[i]); j++ {
			fmt.Print(this.g[i][j], "\t")
		}
		fmt.Println()
	}
}

// 遍历邻边-图算法中最常见的操作
// 需要在外部直接访问图的具体数据g
// 如何保持g本身还是private,又可以在外面遍历具体数据呢
// 使用迭代器可以隐藏迭代的过程,按照一定顺序访问容器内所有的元素
// 实现一个图的迭代器,来访问一个指定节点的所有相邻节点
// 邻边迭代器,传入一个图和一个顶点
// 迭代在这个图中和这个顶点相连的所有顶点
type AdjIteratorSparseGraph struct {
	g     *SparseGraph // 图G的引用
	v     int          // 顶点
	index int
}

// 邻边迭代器的构造函数
func NewAdjIteratorSparseGraph(g Graph, v int) *AdjIteratorSparseGraph {
	if v < 0 || v >= g.V() {
		panic("传入的顶点无效!")
	}
	return &AdjIteratorSparseGraph{g: g.(*SparseGraph), v: v, index: -1}
}

// 返回图G中与顶点v相连接的第一个顶点
func (this *AdjIteratorSparseGraph) Begin() int {
	this.index = 0
	if len(this.g.g[this.v]) > 0 {
		return this.g.g[this.v][this.index]
	}
	// 若没有顶点和v相连接,则返回-1
	return -1
}

// 返回图G中与顶点相连接的下一个顶点
func (this *AdjIteratorSparseGraph) Next() int {
	this.index++
	if this.index < len(this.g.g[this.v]) {
		return this.g.g[this.v][this.index]
	}
	// 若没有顶点和v相连接,则返回-1
	return -1
}

// 查看是否已经迭代完了图G中与顶点v相连接的所有顶点
func (this *AdjIteratorSparseGraph) End() bool {
	return this.index >= len(this.g.g[this.v])
}
