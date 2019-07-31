package _3_Add_Elements_in_BST

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
	if b.root == nil { // tree为空，设置根节点
		b.root = NewNode(e)
		b.size++
	} else {
		b.add(b.root, e)
	}
}

// 向以node为根的二分搜索树中插入元素E，递归算法
func (b *BST) add(n *Node, e interface{}) {
	// 不处理重复数据的节点
	if e == n.e {
		return
	} else if compare(e, n.e) < 0 && n.left == nil {
		n.left = NewNode(e)
		b.size++
		return // 左子树递归终止条件
	} else if compare(e, n.e) > 0 && n.right == nil {
		n.right = NewNode(e)
		b.size++
		return // 右子树递归终止条件
	}

	// 递归调用
	if compare(e, n.e) < 0 {
		b.add(n.left, e)
	} else {
		b.add(n.right, e)
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
