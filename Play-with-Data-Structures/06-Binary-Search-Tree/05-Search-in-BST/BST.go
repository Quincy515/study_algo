package _5_Search_in_BST

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
func (b *BST) add(node *Node, e interface{}) *Node {
	if node == nil {
		b.size++
		return NewNode(e)
	}

	// 递归调用
	if compare(e, node.e) < 0 {
		node.left = b.add(node.left, e)
	} else if compare(e, node.e) > 0 {
		node.right = b.add(node.right, e)
	}
	return node
}

// 看二分搜索树中是否包含元素e
func (b *BST) Contains(e interface{}) bool {
	return b.contains(b.root, e)
}

// 看以node为根的二分搜索树中是否包含元素e，递归算法
func (b *BST) contains(node *Node, e interface{}) bool {
	if node == nil {
		return false
	}

	if compare(e, node.e) == 0 {
		return true
	} else if compare(e, node.e) < 0 {
		return b.contains(node.left, e)
	} else { // compare(e, node.e) > 0
		return b.contains(node.right, e)
	}
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
