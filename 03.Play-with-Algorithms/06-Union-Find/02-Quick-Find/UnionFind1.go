package Quick_Find

// 第一版的Union-Find
type UnionFind struct {
	id    []int // 第一版Union-Find本质就是一个数组
	count int   // 数据个数
}

// 构造函数
func NewUnionFind1(n int) *UnionFind {
	count := n
	id := make([]int, n)
	// 初始化,每一个id[i]指向自己,没有合并的元素
	for i := 0; i < n; i++ {
		id[i] = i // 每个元素都在自己独立的组中,没有互相连接
	}
	return &UnionFind{
		id:    id,
		count: count,
	}
}

// Quick Find 时间复杂度 O(1)
// 查找过程,查找元素p所对应的集合编号
func (uf *UnionFind) Find(p int) int {
	if p < 0 || p >= uf.count {
		panic("p is out of bound.")
	}
	return uf.id[p]
}

// 查看元素p和元素q是否所属一个集合
func (uf UnionFind) IsConnected(p, q int) bool {
	return uf.Find(p) == uf.Find(q)
}

// Quick Find 下的 Union 时间复杂度 O(n)
// 合并元素p和元素q所属的集合
func (uf *UnionFind) UnionElements(p, q int) {
	pID, qID := uf.Find(p), uf.Find(q)

	if pID == qID {
		return
	}

	// 合并过程需要遍历一遍所有元素,将两个元素的所属集合编号合并
	for i := 0; i < uf.count; i++ {
		if uf.id[i] == pID {
			uf.id[i] = qID
		}
	}
}
