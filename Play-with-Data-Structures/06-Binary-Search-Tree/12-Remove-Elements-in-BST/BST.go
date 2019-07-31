package _2_Remove_Elements_in_BST

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

// 二分搜索树的中序遍历
func (b *BST) InOrder() {
	inOrder(b.root)
}

// 中序遍历以node为根的二分搜索树，递归算法
func inOrder(node *Node) {
	if node == nil {
		return
	}

	inOrder(node.left)
	fmt.Println(node.e)
	inOrder(node.right)
}

// 二分搜索树的后序遍历
func (b *BST) PostOrder() {
	postOrder(b.root)
}

// 后序遍历以node为根的二分搜索树，递归算法
func postOrder(node *Node) {
	if node == nil {
		return
	}

	postOrder(node.left)
	postOrder(node.right)
	fmt.Println(node.e)
}

// 寻找二分搜索树的最小元素
func (b *BST) Minimum() interface{} {
	if b.size == 0 {
		panic("BST is empty!")
	}
	return minimum(b.root).e
}

// 返回以node为根的二分搜索树的最小值所在的节点
func minimum(node *Node) *Node {
	if node.left == nil {
		return node
	}
	return minimum(node.left)
}

// 寻找二分搜索树的最大元素
func (b *BST) Maximum() interface{} {
	if b.size == 0 {
		panic("BST is empty!")
	}
	return maximum(b.root).e
}

// 返回以node为根的二分搜索树的最大值所在的节点
func maximum(node *Node) *Node {
	if node.right == nil {
		return node
	}
	return maximum(node.right)
}

// 从二分搜索树中删除最小值所在节点，返回最小值
func (b *BST) RemoveMin() interface{} {
	ret := b.Minimum()
	b.root = b.removeMin(b.root)
	return ret
}

// 删除掉以node为根的二分搜索树中的最小节点
// 返回删除节点后新的二分搜索树的根
func (b *BST) removeMin(node *Node) *Node {
	if node.left == nil {
		rightNode := node.right
		node.right = nil
		b.size--
		return rightNode
	}
	node.left = b.removeMin(node.left)
	return node
}

// 从二分搜索树中删除最大值所在节点
func (b *BST) RemoveMax() interface{} {
	ret := b.Maximum()
	b.root = b.removeMax(b.root)
	return ret
}

// 删除掉以node为根的二分搜索树中的最大节点
// 返回删除节点后新的二分搜索树的根
func (b *BST) removeMax(node *Node) *Node {
	if node.right == nil {
		leftNode := node.left
		node.left = nil
		b.size--
		return leftNode
	}

	node.right = b.removeMax(node.right)
	return node
}

// 从二分搜索树中删除元素为e的节点
func (b *BST) Remove(e interface{}) {
	b.root = b.remove(b.root, e)
}

// 删除掉以node为根的二分搜索树中值为e的节点，递归算法
// 返回删除节点后新的二分搜索树的根
func (b *BST) remove(node *Node, e interface{}) *Node {
	if node == nil {
		return nil
	}

	if compare(e, node.e) < 0 {
		node.left = b.remove(node.left, e)
		return node
	} else if compare(e, node.e) > 0 {
		node.right = b.remove(node.right, e)
		return node
	} else { // compare(e, node.e)=0
		// 待删除节点左子树为空的情况
		if node.left == nil {
			rightNode := node.right
			node.right = nil
			b.size--
			return rightNode
		}
		// 待删除节点右子树为空的情况
		if node.right == nil {
			leftNode := node.left
			node.left = nil
			b.size--
			return leftNode
		}
		// 待删除节点左右子树均不为空的情况
		// 找到比待删除节点大的最小节点，即待删除节点右子树的最小节点
		// 用这个节点顶替待删除节点的位置
		successor := minimum(node.right) // 待删除节点的后继
		successor.right = b.removeMin(node.right)
		successor.left = node.left // 后继节点就顶替了node节点

		node.left, node.right = nil, nil
		return successor
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
