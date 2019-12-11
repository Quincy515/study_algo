package Kruskal

import "strconv"

// Kruskal算法求最小生成树
type KruskalMST struct {
	mst       []*Edge //最小生成树所包含的所有边
	mstWeight weight  // 最小生成树的权值
}

// 构造函数,使用Kruskal算法计算graph的最小生成树
func NewKruskalMST(graph Graph) *KruskalMST {
	k := &KruskalMST{}
	k.mst = make([]*Edge, 0)
	// 使用堆排序,并借助最小堆,最后把边按照权值从小到大取出来
	pq := NewMinHeap(graph.E())      // 将图中所有边存放到一个最小堆中
	for i := 0; i < graph.V(); i++ { // 遍历图中所有的边
		adj := NewAdjIteratorSparseGraph(graph, i)         // 对于每个顶点找其所有的邻边
		for e := adj.Begin(); !adj.End(); e = adj.Next() { // 使用邻边迭代器
			if e.V() < e.W() { // 避免一条边存两次
				pq.Insert(e) // 最小生成树是无向图的
			}
		}
	}

	// 创建一个并查集,来查看已经访问的节点的联通情况
	uf := NewUnionFind(graph.V()) // 空间是顶点个数
	for !pq.IsEmpty() && (len(k.mst) < graph.V()-1) {
		// 从最小堆中依次从小到大取出所有的边
		e := pq.ExtractMin().(*Edge)
		// 如果该边的两个端点是联通的,说明加入这条边将产生环,扔掉这条边
		if uf.IsConnected(e.V(), e.W()) {
			continue
		}
		// 否则,将这条边添加进最小生成树,同时标记边的两个端点联通
		k.mst = append(k.mst, e)
		uf.UnionElements(e.V(), e.W())
	}

	// 计算最小生成树的权值
	k.mstWeight = k.mst[0].Wt()
	for i := 1; i < len(k.mst); i++ {
		aTemp, err := strconv.ParseFloat(k.mstWeight.(string), 32)
		if err != nil {
			panic(err)
		}
		bTemp, err := strconv.ParseFloat(k.mst[i].Wt().(string), 32)
		if err != nil {
			panic(err)
		}
		k.mstWeight = strconv.FormatFloat(aTemp+bTemp, 'f', -1, 32)
	}
	return k
}

// 返回最小生成树的所有边
func (this KruskalMST) mstEdges() []*Edge {
	return this.mst
}

// 返回最小生成树的权值
func (this KruskalMST) result() weight {
	return this.mstWeight
}
