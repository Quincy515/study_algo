package Quick_Union

// 第二版的Union-Find,使用一个数组构建一棵指向父节点的树
type UnionFind struct {
	parent []int // parent[i]表示第i个元素所指向的父节点
	count  int   // 数据个数
}

// 构造函数
func NewUnionFind2(count int) *UnionFind {
	parent := make([]int, count)
	// 初始化,每一个parent[i]指向自己,表示每一个元素自己自成一个集合
	for i := 0; i < count; i++ {
		parent[i] = i
	}
	return &UnionFind{
		parent: parent,
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
	// 根节点的特点: parent[p] == p
	for p != uf.parent[p] {
		p = uf.parent[p]
	}
	return p
}

// 查看元素p和元素q是否所属一个集合
// O(h)复杂度,h为数的高度
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
	uf.parent[pRoot] = qRoot
}
