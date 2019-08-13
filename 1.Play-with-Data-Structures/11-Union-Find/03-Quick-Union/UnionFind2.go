package UnionFind2

// 第二版 Union-Find, 使用一个数组构建一棵指向父节点的树
// parent[i] 表示第一个元素所指向的父节点
type UnionFind2 struct {
	parent []int
}

func NewUnionFind2(size int) *UnionFind2 {
	parent := make([]int, size)
	// 初始化, 每一个 parent[i] 指向自己，表示每一个元素自己自成一个集合
	for i := 0; i < len(parent); i++ {
		parent[i] = i
	}
	return &UnionFind2{parent}
}

func (uf *UnionFind2) GetSize() int {
	return len(uf.parent)
}

// 查找过程，查找元素p所对应的集合编号
//O(h)复杂度，h为树的高度
func (uf *UnionFind2) find(p int) int {
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

// 查询元素p和元素q是否属一个集合
// O(h)复杂度，h为树的高度
func (uf *UnionFind2) IsConnected(p, q int) bool {
	return uf.find(p) == uf.find(q)
}

// 合并元素p和元素q所属的集合
// O(h)复杂度，h为树的高度
func (uf *UnionFind2) UnionElements(p, q int) {
	pRoot := uf.find(p)
	qRoot := uf.find(q)

	if pRoot == qRoot {
		return
	}
	uf.parent[pRoot] = qRoot
}
