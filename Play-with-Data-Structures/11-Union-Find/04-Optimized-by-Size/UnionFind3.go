package UnionFind3

// 第三版Union-Find
type UnionFind struct {
	parent []int // parent[i]表示第一个元素所指向的父节点
	sz     []int // sz[i] 表示以i为根的集合中元素个数
}

func NewUnionFind(size int) *UnionFind {
	sz := make([]int, size)
	parent := make([]int, size)
	// 初始化每一个parent[i]指向自己，表示每一个元素自己自成一个集合
	for i := 0; i < len(parent); i++ {
		sz[i] = i
		parent[i] = i
	}

	return &UnionFind{
		parent: parent,
		sz:     sz,
	}
}

func (uf *UnionFind) GetSize() int {
	return len(uf.parent)
}

// 查找过程，查找元素p所对应的集合编号
// O(h)复杂度，h为树的高度
func (uf *UnionFind) find(p int) int {
	if p < 0 || p > len(uf.parent) {
		panic("p is out of range.")
	}

	// 不断去查询自己的父亲节点，直到到达根节点
	// 根节点的特点: parent[p] == p
	for p != uf.parent[p] {
		p = uf.parent[p]
	}

	return p
}

// 查看元素p和元素q是否所属一个集合
// O(h)复杂度，h为树的高度
func (uf *UnionFind) IsConnected(p, q int) bool {
	return uf.find(p) == uf.find(q)
}

// 合并元素p和元素q所属的集合
// O(h)复杂度，h为树的高度
func (uf *UnionFind) UnionElements(p, q int) {
	pRoot, qRoot := uf.find(p), uf.find(q)

	if pRoot == qRoot {
		return
	}

	// 根据两个元素所在树的元素个数不同判断合并方向
	// 将元素个数少的集合合并到元素个数多的集合上
	if uf.sz[pRoot] < uf.sz[qRoot] {
		uf.parent[pRoot] = qRoot
		uf.sz[qRoot] += uf.sz[pRoot]
	} else { // sz[qRoot] <= sz[pRoot]
		uf.parent[qRoot] = pRoot
		uf.sz[pRoot] += uf.sz[qRoot]
	}
}
