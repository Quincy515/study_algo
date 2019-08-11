package RBTree

import (
	"fmt"
	"reflect"
)

const RED = true
const BLACK = false

type Node struct {
	key, value  interface{}
	left, right *Node
	color       bool // True代表红色，False代表黑色
}

// 生成Node节点
func NewNode(key, value interface{}) *Node {
	return &Node{
		key:   key,
		value: value,
		left:  nil,
		right: nil,
		color: RED, // 默认为融合3-节点
	}
}

type RBTree struct {
	root *Node
	size int
}

func NewRBTree() *RBTree {
	return &RBTree{}
}

func (rb *RBTree) GetSize() int {
	return rb.size
}

func (rb *RBTree) IsEmpty() bool {
	return rb.size == 0
}

// 判断节点node的颜色
func isRed(node *Node) bool {
	if node == nil { // 空节点为黑色
		return BLACK
	}
	return node.color
}

// 向二分搜索树中添加元素(key, value)
func (rb *RBTree) Add(key interface{}, value interface{}) {
	rb.root = rb.add(rb.root, key, value)
}

// 向以node为根的二分搜索树中插入元素(key, value),递归算法
// 返回插入新节点后二分搜索树的根
func (rb *RBTree) add(node *Node, key interface{}, value interface{}) *Node {
	if node == nil {
		rb.size++
		return NewNode(key, value)
	}
	if compare(key, node.key) < 0 {
		node.left = rb.add(node.left, key, value)
	} else if compare(key, node.key) > 0 {
		node.right = rb.add(node.right, key, value)
	} else { // compare(key, node.key) == 0
		node.value = value
	}
	return node
}

// 返回以node为根节点的二分搜索树中，key所在的节点
func (rb *RBTree) getNode(node *Node, key interface{}) *Node {
	if node == nil { // 未找到等于key的节点
		return nil
	}
	if compare(key, node.key) == 0 {
		return node
	} else if compare(key, node.key) < 0 {
		return rb.getNode(node.left, key)
	} else {
		return rb.getNode(node.right, key)
	}
}

func (rb *RBTree) Contains(key interface{}) bool {
	return rb.getNode(rb.root, key) != nil
}

func (rb *RBTree) Get(key interface{}) interface{} {
	node := rb.getNode(rb.root, key)
	if node == nil {
		return nil
	} else {
		return node.value
	}
}

func (rb *RBTree) Set(key interface{}, newValue interface{}) {
	node := rb.getNode(rb.root, key)
	if node == nil {
		panic(fmt.Sprintf("%v, doesn't exist.", key))
	}
	node.value = newValue
}

// 从二分搜索树中删除键为key的节点
func (rb *RBTree) Remove(key interface{}) interface{} {
	node := rb.getNode(rb.root, key)
	if node != nil {
		rb.root = rb.remove(rb.root, key)
		return node.value
	}

	return nil
}

// 删除掉以node为根的二分搜索树中键为key的节点，递归算法
// 返回删除节点后新的二分搜索树的根
func (rb *RBTree) remove(node *Node, key interface{}) *Node {
	if node == nil {
		return nil
	}
	if compare(key, node.key) < 0 {
		node.left = rb.remove(node.left, key)
		return node
	} else if compare(key, node.key) > 0 {
		node.right = rb.remove(node.right, key)
		return node
	} else { // compare(key, node.key)==0 删除当前节点
		// 待删除节点左子树为空的情况
		if node.left == nil {
			rightNode := node.right
			node.right = nil
			rb.size--
			return rightNode
		}
		// 待删除节点右子树为空的情况
		if node.right == nil {
			leftNode := node.left
			node.left = nil
			rb.size--
			return leftNode
		}
		// 待删除节点左右子树均不为空的情况
		// 找到比待删除节点大的最小节点，即待删除节点右子树的最小节点
		// 用这个节点顶替待删除节点的位置
		successor := rb.minimum(node.right)
		successor.right = rb.removeMin(node.right)
		successor.left = node.left

		node.left, node.right = nil, nil
		return successor
	}
}

// 返回以node为根的二分搜索树的最小值所在的节点
func (rb *RBTree) minimum(node *Node) *Node {
	if node.left == nil {
		return node
	}
	return rb.minimum(node.left)
}

// 删除掉以node为根的二分搜索树中的最小节点
// 返回删除节点后新的二分搜索树的根
func (rb *RBTree) removeMin(node *Node) *Node {
	if node.left == nil {
		rightNode := node.right
		node.right = nil
		rb.size--

		return rightNode
	}

	node.left = rb.removeMin(node.left)
	return node
}

func compare(a, b interface{}) int {
	aType := reflect.TypeOf(a).String()
	bType := reflect.TypeOf(b).String()
	if aType != bType {
		panic("cannot compare different type params")
	}
	switch a.(type) {
	case int:
		if a.(int) > b.(int) {
			return 1
		} else if a.(int) < b.(int) {
			return -1
		} else {
			return 0
		}
	case string:
		if a.(string) > b.(string) {
			return 1
		} else if a.(string) < b.(string) {
			return -1
		} else {
			return 0
		}
	default:
		panic("unsupported type params")
	}
}
