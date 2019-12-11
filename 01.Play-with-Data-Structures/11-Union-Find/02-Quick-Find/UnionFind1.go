package UnionFind1

// 第一版Union-find
type UnionFind struct {
	id []int
}

func NewUnionFind(size int) *UnionFind {
	id := make([]int, size)
	// 初始化，每一个id[i]指向自己，没有合并的元素
	for i := 0; i < size; i++ {
		id[i] = i
	}
	return &UnionFind{id}
}

func (uf *UnionFind) GetSize() int {
	return len(uf.id)
}

// 查找元素p所对应的集合编号
func (uf *UnionFind) find(p int) int {
	if p < 0 || p >= len(uf.id) {
		panic("p is out of bound.")
	}
	return uf.id[p]
}

// 查看元素p和元素q是否属于一个集合
func (uf *UnionFind) IsConnected(p, q int) bool {
	return uf.find(p) == uf.find(q)
}

// 合并元素p和元素q所属的集合
func (uf *UnionFind) UnionElements(p, q int) {
	pID, qID := uf.find(p), uf.find(q)
	if pID == qID {
		return
	}

	for i := 0; i < len(uf.id); i++ {
		if uf.id[i] == pID {
			uf.id[i] = qID
		}
	}
}
