package Optimize_by_Size

// 第三版Union-Find
type UnionFind struct {
	parent []int // parent[i] 表示第i个元素所指向的父节点
	sz     []int // sz[i] 表示以i为根的集合中的元素个数
	count  int   // 数据个数
}

// 构造函数
func NewUnionFind(count int) *UnionFind {
	parent := make([]int, count)
	sz := make([]int, count)
	for i := 0; i < count; i++ {
		parent[i] = i
		sz[i] = 1
	}
	return &UnionFind{
		parent: parent,
		sz:     sz,
		count:  count,
	}
}

// 查找过程,查找元素p所对应的集合编号
// O(h)复杂度,h为树的高度
func (uf *UnionFind) Find(p int) int {
	if p < 0 || p >= uf.count {
		panic("p is out of range.")
	}
	// 不断去查询自己的父节点,直到到达根节点
	for p != uf.parent[p] {
		p = uf.parent[p]
	}
	return p
}

// 查看元素p和元素q是否所属一个集合
// O(h)复杂度,h为树的高度
func (uf *UnionFind) IsConnected(p, q int) bool {
	return uf.Find(p) == uf.Find(q)
}

// 合并元素p和元素q所属的集合
// O(h)复杂度,h为树的高度
func (uf *UnionFind) UnionElements(p, q int) {
	pRoot, qRoot := uf.Find(p), uf.Find(q)
	if pRoot == qRoot {
		return
	}

	// 根据两个元素所在树的元素个数不同判断合并方向
	// 将元素个数少的集合合并到元素个数多的集合上
	if uf.sz[pRoot] < uf.sz[qRoot] {
		uf.parent[pRoot] = qRoot
		uf.sz[qRoot] += uf.sz[pRoot]
	} else {
		uf.parent[qRoot] = pRoot
		uf.sz[pRoot] += uf.sz[qRoot]
	}
}
