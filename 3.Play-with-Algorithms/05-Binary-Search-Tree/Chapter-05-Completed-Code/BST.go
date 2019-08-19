package Binary_Search_Tree

import (
	"fmt"
)

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

// 二分搜索树的层序遍历
func (bst *BST) LevelOrder() {
	var q []*node           // 申请一个队列
	q = append(q, bst.root) // 入队进入队尾
	for len(q) > 0 {        // 队列不为空
		n := q[0] // 从队首出队
		q = q[1:] //删除出队元素
		fmt.Print(n.key, " ")

		if n.left != nil { // 左子树不为空进行入队操作
			q = append(q, n.left)
		}
		if n.right != nil { // 右子树不为空进行入队操作
			q = append(q, n.right)
		}
	}
}

// 寻找最小的键值
func (bst *BST) Minimum() int {
	if bst.count == 0 {
		panic("二分搜索树中没有元素!")
	}

	minNode := minimum(bst.root)
	return minNode.key
}

// 返回以node为根的二分搜索树的最小键值所在的节点
func minimum(n *node) *node {
	if n.left == nil {
		return n
	}
	return minimum(n.left)
}

// 寻找二分搜索树的最大的键值
func (bst *BST) Maximum() int {
	if bst.count == 0 {
		panic("二分搜索树中没有元素!")
	}
	maxNode := maximum(bst.root)
	return maxNode.key
}

// 返回以node为根的二分搜索树的最大键值所在的节点
func maximum(n *node) *node {
	if n.right == nil {
		return n
	}
	return maximum(n.right)
}

// 从二分搜索树中删除最小值所在节点
func (bst *BST) RemoveMin() {
	if bst.root != nil {
		bst.root = bst.removeMin(bst.root)
	}
}

// 删除掉以node为根的二分搜索树中的最小节点
// 返回删除节点后新的二分搜索树的根
func (bst *BST) removeMin(n *node) *node {
	if n.left == nil {
		rightNode := n.right
		n.right = nil
		bst.count--
		return rightNode
	}
	n.left = bst.removeMin(n.left)
	return n
}

// 从二分搜索树中删除最大值所在节点
func (bst *BST) RemoveMax() {
	if bst.root != nil {
		bst.root = bst.removeMax(bst.root)
	}
}

// 删除掉以node为根的二分搜索树中的最大节点
// 返回删除节点后新的二分搜索树的根
func (bst *BST) removeMax(n *node) *node {
	if n.right == nil {
		leftNode := n.left
		n.left = nil
		bst.count--
		return leftNode
	}
	n.right = bst.removeMax(n.right)
	return n
}

// 从二分搜索树中删除键值为key的节点
func (bst *BST) Remove(key int) {
	bst.root = bst.remove(bst.root, key)
}

// 删除掉以node为根的二分搜索树中键值为key的节点,递归算法
// 返回删除节点后新的二分搜索树的根
func (bst *BST) remove(n *node, key int) *node {
	if n == nil {
		return nil
	}

	if key < n.key {
		n.left = bst.remove(n.left, key)
		return n
	} else if key > n.key {
		n.right = bst.remove(n.right, key)
		return n
	} else { // key == n.key

		// 待删除节点左子树为空的情况
		if n.left == nil {
			rightNode := n.right
			n.right = nil
			bst.count--
			return rightNode
		}

		// 待删除节点右子树为空的情况
		if n.right == nil {
			leftNode := n.left
			n.left = nil
			bst.count--
			return leftNode
		}

		// 待删除节点左右子树均不为空的情况

		// 找到比待删除节点大的最小节点,即待删除节点右子树的最小节点
		// 用这个节点顶替待删除节点的位置
		successor := minimum(n.right)
		successor.right = bst.removeMin(n.right)
		successor.left = n.left

		n.left, n.right = nil, nil
		return successor
	}
}
