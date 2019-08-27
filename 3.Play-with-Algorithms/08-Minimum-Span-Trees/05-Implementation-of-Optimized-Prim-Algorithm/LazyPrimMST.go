package Prim

import (
	"strconv"
)

// 使用Prim算法求图的最小生成树
type LazyPrimMST struct {
	g         Graph    // 图的引用
	pq        *MinHeap // 最小堆,算法辅助数据结构
	marked    []bool   // 标记数组,在算法运行过程中标记节点i是否被访问
	mst       []*Edge  // 最小生成树所包含的所有边
	mstWeight weight   // 最小生成树的权值
}

// 构造函数,使用Prim算法求图的最小生成树
func NewLazyPrimMST(graph Graph) *LazyPrimMST {
	p := &LazyPrimMST{}
	// 算法初始化
	p.g = graph
	p.pq = NewMinHeap(graph.E())
	p.marked = make([]bool, graph.V())
	p.mst = make([]*Edge, 0)

	// Lazy Prim
	p.visit(0)
	for !p.pq.IsEmpty() {
		//fmt.Print(" 使用最小堆找出已经访问的边中权值最小的边: ", p.pq.ExtractMin(), " ")
		// 使用最小堆找出已经访问的边中权值最小的边
		e := p.pq.ExtractMin().(*Edge)
		// 如果这条边的两端都已经访问过了,则扔掉这条边
		if p.marked[e.V()] == p.marked[e.W()] {
			continue
		}
		// 否则,这条边则应该存储在最小生成树中
		p.mst = append(p.mst, e)

		// 访问和这条边连接的还没有被访问过的节点
		if !p.marked[e.V()] {
			p.visit(e.V())
		} else {
			p.visit(e.W())
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
	}

	return p
}

// 辅助函数,范文节点v
func (this *LazyPrimMST) visit(v int) {
	assert(!this.marked[v])
	this.marked[v] = true

	// 将和节点v相连接的所有未访问的边放入最小堆中
	adj := NewAdjIteratorSparseGraph(this.g, v)
	for e := adj.Begin(); !adj.End(); e = adj.Next() {
		if !this.marked[e.Other(v)] {
			this.pq.Insert(e)
		}
	}
}

// 返回最小生成树的所有边
func (this *LazyPrimMST) mstEdges() []*Edge {
	return this.mst
}

// 返回最小生成树的权值
func (this LazyPrimMST) result() weight {
	return this.mstWeight
}
