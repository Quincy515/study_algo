package AVLTree

import (
	"fmt"
	"math"
	"reflect"
)

type Node struct {
	key         interface{}
	value       interface{}
	left, right *Node
	height      int // 标注节点的高度
}

type AVLTree struct {
	root *Node
	size int
}

func NewAVLTree() *AVLTree {
	return &AVLTree{&Node{
		key:    "",
		value:  0,
		left:   nil,
		right:  nil,
		height: 1,
	}, 0}
}

func (avl *AVLTree) GetSize() int {
	return avl.size
}

func (avl *AVLTree) IsEmpty() bool {
	return avl.size == 0
}

// 获得节点Node的高度
func (avl *AVLTree) getHeight(node *Node) int {
	if node == nil {
		return 0
	} else {
		return node.height
	}
}

// 获得节点Node的平衡因子
func (avl *AVLTree) getBalanceFactor(node *Node) int {
	if node == nil {
		return 0
	}
	return avl.getHeight(node.left) - avl.getHeight(node.right)
}

// 向二分搜索树中添加元素(key, value)
func (avl *AVLTree) Add(key interface{}, value interface{}) {
	avl.root = avl.add(avl.root, key, value)
}

// 向以node为根的二分搜索树中插入元素(key, value),递归算法
// 返回插入新节点后二分搜索树的根
func (avl *AVLTree) add(node *Node, key interface{}, value interface{}) *Node {
	if node == nil {
		avl.size++
		return &Node{
			key:    key,
			value:  value,
			left:   nil,
			right:  nil,
			height: 1,
		}
	}
	if compare(key, node.key) < 0 {
		node.left = avl.add(node.left, key, value)
	} else if compare(key, node.key) > 0 {
		node.right = avl.add(node.right, key, value)
	} else { // compare(key, node.key) == 0
		node.value = value
	}

	// 更新height
	node.height = 1 + int(math.Max(
		float64(avl.getHeight(node.left)),
		float64(avl.getHeight(node.right))))
	// 计算平衡因子
	balanceFactor := avl.getBalanceFactor(node)
	if math.Abs(float64(balanceFactor)) > 1 {
		fmt.Println("unbalanced: ", balanceFactor)
	}
	return node
}

// 返回以node为根节点的二分搜索树中，key所在的节点
func (avl *AVLTree) getNode(node *Node, key interface{}) *Node {
	if node == nil { // 未找到等于key的节点
		return nil
	}
	if compare(key, node.key) == 0 {
		return node
	} else if compare(key, node.key) < 0 {
		return avl.getNode(node.left, key)
	} else {
		return avl.getNode(node.right, key)
	}
}

func (avl *AVLTree) Contains(key interface{}) bool {
	return avl.getNode(avl.root, key) != nil
}

func (avl *AVLTree) Get(key interface{}) interface{} {
	node := avl.getNode(avl.root, key)
	if node == nil {
		return nil
	} else {
		return node.value
	}
}

func (avl *AVLTree) Set(key interface{}, newValue interface{}) {
	node := avl.getNode(avl.root, key)
	if node == nil {
		panic(fmt.Sprintf("%v, doesn't exist.", key))
	}
	node.value = newValue
}

// 从二分搜索树中删除键为key的节点
func (avl *AVLTree) Remove(key interface{}) interface{} {
	node := avl.getNode(avl.root, key)
	if node != nil {
		avl.root = avl.remove(avl.root, key)
		return node.value
	}

	return nil
}

// 删除掉以node为根的二分搜索树中键为key的节点，递归算法
// 返回删除节点后新的二分搜索树的根
func (avl *AVLTree) remove(node *Node, key interface{}) *Node {
	if node == nil {
		return nil
	}
	if compare(key, node.key) < 0 {
		node.left = avl.remove(node.left, key)
		return node
	} else if compare(key, node.key) > 0 {
		node.right = avl.remove(node.right, key)
		return node
	} else { // compare(key, node.key)==0 删除当前节点
		// 待删除节点左子树为空的情况
		if node.left == nil {
			rightNode := node.right
			node.right = nil
			avl.size--
			return rightNode
		}
		// 待删除节点右子树为空的情况
		if node.right == nil {
			leftNode := node.left
			node.left = nil
			avl.size--
			return leftNode
		}
		// 待删除节点左右子树均不为空的情况
		// 找到比待删除节点大的最小节点，即待删除节点右子树的最小节点
		// 用这个节点顶替待删除节点的位置
		successor := avl.minimum(node.right)
		successor.right = avl.removeMin(node.right)
		successor.left = node.left

		node.left, node.right = nil, nil
		return successor
	}
}

// 返回以node为根的二分搜索树的最小值所在的节点
func (avl *AVLTree) minimum(node *Node) *Node {
	if node.left == nil {
		return node
	}
	return avl.minimum(node.left)
}

// 删除掉以node为根的二分搜索树中的最小节点
// 返回删除节点后新的二分搜索树的根
func (avl *AVLTree) removeMin(node *Node) *Node {
	if node.left == nil {
		rightNode := node.right
		node.right = nil
		avl.size--

		return rightNode
	}

	node.left = avl.removeMin(node.left)
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
