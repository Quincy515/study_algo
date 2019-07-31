package _6_PreOrder_Traverse_in_BST

import (
	"bytes"
	"fmt"
)

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
	return contains(b.root, e)
}

// 看以node为根的二分搜索树中是否包含元素e，递归算法
func contains(node *Node, e interface{}) bool {
	if node == nil {
		return false
	}

	if compare(e, node.e) == 0 {
		return true
	} else if compare(e, node.e) < 0 {
		return contains(node.left, e)
	} else { // compare(e, node.e) > 0
		return contains(node.right, e)
	}
}

// 二分搜索树的前序遍历
func (b *BST) PreOrder() {
	preOrder(b.root)
}

// 前序遍历以node为根的二分搜索树，递归算法
func preOrder(node *Node) {
	if node == nil {
		return
	}
	//if node != nil {
	fmt.Println(node.e)
	preOrder(node.left)
	preOrder(node.right)
	//}
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

func (b *BST) String() string {
	var buffer bytes.Buffer
	generateBSTString(b.root, 0, &buffer)
	return buffer.String()
}

// 生成以Node为根节点，深度为depth的描述二叉树的字符串
func generateBSTString(node *Node, depth int, buffer *bytes.Buffer) {
	if node == nil {
		buffer.WriteString(generateDepthString(depth) + "nil\n")
		return
	}
	buffer.WriteString(generateDepthString(depth) + fmt.Sprint(node.e) + "\n")
	generateBSTString(node.left, depth+1, buffer)
	generateBSTString(node.right, depth+1, buffer)
}

func generateDepthString(depth int) string {
	var buffer bytes.Buffer
	for i := 0; i < depth; i++ {
		buffer.WriteString("--")
	}
	return buffer.String()
}
