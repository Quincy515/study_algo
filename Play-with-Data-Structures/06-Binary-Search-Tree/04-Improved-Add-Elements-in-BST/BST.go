package _4_Improved_Add_Elements_in_BST

type Node struct {
	e           interface{}
	left, right *Node
}

// 构造函数
func NewNode(e interface{}) *Node {
	return &Node{e: e}
}

type BST struct {
	root *Node
	size int
}

// 二分搜索树的构造函数
func NewBST() *BST {
	return &BST{}
}

func (b *BST) Size() int {
	return b.size
}

func (b *BST) IsEmpty() bool {
	return b.size == 0
}

// 向二分搜索树中添加新的元素e
func (b *BST) Add(e interface{}) {
	b.root = b.add(b.root, e)
}

// 向以node为根的二分搜索树中插入元素E，递归算法
func (b *BST) add(n *Node, e interface{}) *Node {
	if n == nil {
		b.size++
		return NewNode(e)
	}

	// 递归调用
	if compare(e, n.e) < 0 {
		n.left = b.add(n.left, e)
	} else if compare(e, n.e) > 0 {
		n.right = b.add(n.right, e)
	}
	return n
}

func compare(a, b interface{}) int {
	if a.(int) > b.(int) {
		return 1
	} else if a.(int) < b.(int) {
		return -1
	} else {
		return 0
	}
}
