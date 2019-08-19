package Path_Compression2

// 第五版Union-Find
// http://coding.imooc.com/learn/questiondetail/7287.html
// 在后序的代码中,并不会维护rank的语意,也就是rank的值在路径压缩的过程中,有可能不再是树的层数值
// 这也是rank不叫height或depth的原因,他只是作为比较的一个标准
type UnionFind struct {
	parent []int // parent[i] 表示第i个元素所指向的父节点
	rank   []int // rank[i] 表示以i为根的集合所代表的层数
	count  int   // 数据个数
}

// 构造函数
func NewUnionFind(count int) *UnionFind {
	parent := make([]int, count)
	rank := make([]int, count)
	for i := 0; i < count; i++ {
		parent[i] = i
		rank[i] = 1
	}
	return &UnionFind{
		parent: parent,
		rank:   rank,
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
	//for p != uf.parent[p] {
	//	// path compression路径压缩
	//	uf.parent[p] = uf.parent[uf.parent[p]]
	//	p = uf.parent[p]
	//}
	//return p

	//path compression路径压缩递归算法
	if p != uf.parent[p] {
		uf.parent[p] = uf.Find(uf.parent[p])
	}
	return uf.parent[p]
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
	if uf.rank[pRoot] < uf.rank[qRoot] {
		uf.parent[pRoot] = qRoot
	} else if uf.rank[qRoot] < uf.rank[pRoot] {
		uf.parent[qRoot] = pRoot
	} else { // rank[pRoot] == rank[qRoot]
		uf.parent[pRoot] = qRoot
		uf.rank[qRoot] += 1 // 此时,维护rank的值
	}
}
