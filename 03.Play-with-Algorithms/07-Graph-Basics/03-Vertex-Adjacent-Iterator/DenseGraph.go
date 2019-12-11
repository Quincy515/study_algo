package Graph_Representation

// 稠密图-邻接矩阵
type DenseGraph struct {
	n, m     int      // 节点数和边数
	directed bool     // 是否为有向图
	g        [][]bool // 图的具体数据
}

// 构造函数
func NewDenseGraph(n int, directed bool) *DenseGraph {
	if n < 0 {
		panic("节点数必须大于0")
	}
	// g 初始化为n*n的布尔矩阵,每一个g[i][j]均为false,表示没有任何边
	g := make([][]bool, n)
	for i := 0; i < n; i++ {
		g[i] = make([]bool, n) // false为boolean型变量的默认值
	}

	return &DenseGraph{
		n:        n,        // 节点个数
		m:        0,        // 初始化没有任何边
		directed: directed, // 是否为有向图
		g:        g,        // 图的具体数据
	}
}

// 返回节点个数
func (this *DenseGraph) V() int {
	return this.n
}

// 返回边的个数
func (this *DenseGraph) E() int {
	return this.m
}

// 向图中添加一条边
func (this *DenseGraph) AddEdge(v, w int) {
	if v < 0 || v >= this.n || w < 0 || w >= this.n {
		panic("向图中添加一条边超出了图的最大范围")
	}

	if this.hashEdge(v, w) {
		return // v 和 w 之间本身有边了
	}
	this.g[v][w] = true // 记录v和w是有边的
	if !this.directed { // 如果是无向图,相应(w,v)的位置也要为true
		this.g[w][v] = true
	}
	this.m++ // 添加边的数目
}

// 验证图中是否有从v到w的边
func (this *DenseGraph) hashEdge(v, w int) bool {
	if v < 0 || v >= this.n || w < 0 || w >= this.n {
		panic("向图中添加一条边超出了图的最大范围")
	}
	return this.g[v][w]
}

// 邻边迭代器,传入一个图和一个顶点
// 迭代在这个图中和这个顶点相连的所有顶点
type AdjIteratorDenseGraph struct {
	g        *DenseGraph // 图G的引用
	v, index int
}

// 邻边迭代器构造函数
func NewAdjIteratorDenseGraph(g *DenseGraph, v int) *AdjIteratorDenseGraph {
	if v < 0 || v >= g.n {
		panic("传入的顶点v越界!")
	}
	return &AdjIteratorDenseGraph{g: g, v: v, index: -1}
}

// 返回图G中与顶点v相连接的第一个顶点
func (this *AdjIteratorDenseGraph) Begin() int {
	// 索引从-1开始,因为每次遍历都需要调用一次Next()
	this.index = -1
	return this.Next()
}

// 返回图G中顶点v相连接的下一个顶点
func (this *AdjIteratorDenseGraph) Next() int {
	// 从当前index开始向后搜索,直到找到一个g[v][index]为true
	for this.index += 1; this.index < this.g.V(); this.index++ {
		if this.g.g[this.v][this.index] {
			return this.index
		}
	}
	// 若没有顶点和v相连接,则返回-1
	return -1
}

// 查看是否已经迭代完了图G中与顶点v相连的所有顶点
func (this *AdjIteratorDenseGraph) End() bool {
	return this.index >= this.g.V()
}
