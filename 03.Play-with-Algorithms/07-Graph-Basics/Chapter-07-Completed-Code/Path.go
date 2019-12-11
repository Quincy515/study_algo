package Graph

import "fmt"

type Path struct {
	g       Graph  // 图的引用
	s       int    // 起始点
	visited []bool // 记录dfs的过程中节点是否被访问
	from    []int  // 记录路径,from[i] 表示查找的路径上i的上一个节点
}

// 构造函数,寻路算法,寻找图graph从s点到其他点的路径
func NewPath(graph Graph, s int) *Path {
	// 算法初始化
	assert(s >= 0 && s < graph.V()) // 起始点(源点)不越界
	p := &Path{}
	p.g = graph
	p.s = s // 起始点(源点)
	p.visited = make([]bool, graph.V())
	p.from = make([]int, graph.V())
	for i := 0; i < graph.V(); i++ {
		p.visited[i] = false
		p.from[i] = -1
	}

	// 寻路算法
	p.dfs(s)
	return p
}

// 图的深度优先遍历
func (this *Path) dfs(v int) {
	this.visited[v] = true
	adj := NewAdjIteratorSparseGraph(this.g, v)
	for i := adj.Begin(); !adj.End(); i = adj.Next() {
		if !this.visited[i] {
			this.from[i] = v
			this.dfs(i)
		}
	}
}

// 查看从s点到w点是否有路径
func (this *Path) hashPath(w int) bool {
	assert(w >= 0 && w < this.g.V())
	return this.visited[w]
}

// 查询从s点到w点的路径,存放在vec中
func (this *Path) path(w int, vec *[]int) {
	assert(this.hashPath(w))
	var s []int
	// 通过from数组逆向查找到从s到w的路径,存放到栈中
	p := w
	for p != -1 { // 将路径以倒序的方式存入到stack中
		s = append(s, p)
		p = this.from[p]
	}
	// 清空vec数组
	//for i := 0; i < len(*vec); i++ {
	//	*vec[i] = -1
	//}

	// 从栈中依次取出元素,获得顺序的从s到w的路径
	for len(s) > 0 {
		*vec = append(*vec, s[len(s)-1]) // 栈顶元素
		s = s[:len(s)-1]                 // 出栈
	}
}

// 打印出从s点到w点的路径
func (this *Path) showPath(w int) {
	assert(this.hashPath(w))
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
