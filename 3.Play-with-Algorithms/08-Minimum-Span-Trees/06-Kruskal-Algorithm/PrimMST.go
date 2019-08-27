package Kruskal

import "strconv"

// 使用优化的Prim算法求图的最小生成树
type PrimMST struct {
	g         Graph         // 图的引用
	ipq       *IndexMinHeap // 最小索引堆,算法辅助数据结构
	edgeTo    []*Edge       // 访问的点所对应的边,算法辅助数据结构
	marked    []bool        // 标记数组,在算法运行过程中标记节点i是否被访问
	mst       []*Edge       // 最小生成树所包含的所有边
	mstWeight weight        // 最小生成树的权值
}

// 构造函数,使用Prim算法求图的最小生成树
func NewPrimMST(graph Graph) *PrimMST {
	p := &PrimMST{}
	p.g = graph
	assert(graph.E() >= 1)
	p.ipq = NewIndexMinHeap(graph.V())

	// 算法初始化
	p.marked = make([]bool, graph.V())
	p.edgeTo = make([]*Edge, graph.V())
	for i := 0; i < graph.V(); i++ {
		p.marked[i] = false
		p.edgeTo[i] = nil
	}
	p.mst = make([]*Edge, 0)

	// Prim
	p.visit(0)
	for !p.ipq.IsEmpty() {
		// 使用最小索引堆找出已经访问的边中权值最小的边
		// 最小索引堆中存储的是点的索引,通过点的索引找到相对应的边
		v := p.ipq.ExtractMinIndex()
		assert(p.edgeTo[v] != nil)
		p.mst = append(p.mst, p.edgeTo[v])
		p.visit(v)
	}

	// 计算最小生成树的权值
	p.mstWeight = p.mst[0].Wt()
	for i := 1; i < len(p.mst); i++ {
		aTemp, err := strconv.ParseFloat(p.mstWeight.(string), 32)
		if err != nil {
			panic(err)
		}
		bTemp, err := strconv.ParseFloat(p.mst[i].Wt().(string), 32)
		if err != nil {
			panic(err)
		}
		p.mstWeight = strconv.FormatFloat(aTemp+bTemp, 'f', -1, 32)
	}

	return p
}

// 访问节点v
func (this *PrimMST) visit(v int) {
	assert(!this.marked[v])
	this.marked[v] = true

	// 将和节点v相连接的未访问的另一端点,和与之相连接的边,放入最小堆中
	adj := NewAdjIteratorSparseGraph(this.g, v)
	for e := adj.Begin(); !adj.End(); e = adj.Next() {
		w := e.Other(v)      // 返回节点v相连接的另一个节点
		if !this.marked[w] { // 如果边的另一端点未被访问
			// 如果从没有考虑过这个端点,直接将这个端点和与之相连接的边加入索引堆
			if this.edgeTo[w] == nil { // 查看是否之前找到过和w相邻的横切边
				this.ipq.Insert(e.Wt(), w) // 添加索引和权值
				this.edgeTo[w] = e         // edgeTo数组对w这个索引进行更新
			} else if compare(e.Wt(), this.edgeTo[w].Wt()) < 0 {
				// 如果曾经考虑这个端点,但现在的边比之前考虑的边更短,则进行替换
				this.edgeTo[w] = e
				this.ipq.change(w, e.Wt())
			}
		}
	}
}

// 返回最小生成树的所有边
func (this PrimMST) mstEdges() []*Edge {
	return this.mst
}

// 返回最小生成树的权值
func (this PrimMST) result() weight {
	return this.mstWeight
}
