package Graph_Representation

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
