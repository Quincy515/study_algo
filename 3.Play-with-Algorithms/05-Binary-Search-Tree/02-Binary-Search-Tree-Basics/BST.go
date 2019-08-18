package Binary_Search_Tree

// 二分搜索树中的节点为私有的结构体,外界不需要了解二分搜索树节点的具体实现
type node struct {
	key, value  int
	left, right *node
}

type BST struct {
	root  *node // 根节点
	count int   // 节点个数
}

// 构造函数,默认构造一棵空二分搜索树
func NewBst() *BST {
	return &BST{}
}

// 返回二分搜索树的节点个数
func (bst *BST) Size() int {
	return bst.count
}

// 返回二分搜索树是否为空
func (bst *BST) IsEmpty() bool {
	return bst.count == 0
}
