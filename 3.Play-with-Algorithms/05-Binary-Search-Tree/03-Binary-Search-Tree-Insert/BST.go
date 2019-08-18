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

// 向二分搜索树中插入一个新的(key,value)数据对
func (bst *BST) Insert(key, value int) {
	bst.root = bst.insert(bst.root, key, value)
}

// 向以node为根的二叉搜索树中,插入节点(key,value)
// 返回插入新节点后的二叉搜索树的根
func (bst *BST) insert(n *node, key, value int) *node {
	if n == nil {
		bst.count++
		return &node{
			key:   key,
			value: value,
			left:  nil,
			right: nil,
		}
	}

	if key == n.key {
		n.value = value
	} else if key < n.key {
		n.left = bst.insert(n.left, key, value)
	} else { // key > node.left
		n.right = bst.insert(n.right, key, value)
	}
	return n
}
