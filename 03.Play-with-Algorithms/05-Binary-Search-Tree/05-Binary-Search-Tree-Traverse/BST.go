package Binary_Search_Tree

import "fmt"

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

// 查看二分搜索树中是否存在键key
func (bst *BST) Contain(key int) bool {
	return contain(bst.root, key)
}

// 查看以node为根的二分搜索树中是否包含键值为key的节点,使用递归算法
func contain(n *node, key int) bool {
	if n == nil {
		return false
	}

	if key == n.key {
		return true
	} else if key < n.key {
		return contain(n.left, key)
	} else { // key > n.right
		return contain(n.right, key)
	}
}

// 在二分搜索树中搜索键key所对应的值。如果这个值不存在,则返回nil
func (bst *BST) Search(key int) int {
	return search(bst.root, key)
}

// 在以node为根的二分搜索树中查找key所对应的value,递归算法
// 若value不存在,则返回nil
func search(n *node, key int) int {
	if n == nil {
		return 0
	}

	if key == n.key {
		return n.value
	} else if key < n.key {
		return search(n.left, key)
	} else { // key > n.right
		return search(n.right, key)
	}
}

// 二分搜索树额前序遍历
func (bst *BST) PreOrder() {
	preOrder(bst.root)
}

// 对以node为根的二叉搜索树进行前序遍历,递归算法
func preOrder(n *node) {
	if n != nil {
		fmt.Print(n.key, " ")
		preOrder(n.left)
		preOrder(n.right)
	}
}

// 二分搜索树的中序遍历
func (bst *BST) InOrder() {
	inOrder(bst.root)
}

// 对以node为根的二分搜索树进行中序遍历,递归算法
func inOrder(n *node) {
	if n != nil {
		inOrder(n.left)
		fmt.Print(n.key, " ")
		inOrder(n.right)
	}

}

// 二分搜索树的后序遍历
func (bst *BST) PostOrder() {
	postOrder(bst.root)
}

// 对以node为根的二分搜索树进行后序遍历,递归算法
func postOrder(n *node) {
	if n != nil {
		postOrder(n.left)
		postOrder(n.right)
		fmt.Print(n.key, " ")
	}
}
