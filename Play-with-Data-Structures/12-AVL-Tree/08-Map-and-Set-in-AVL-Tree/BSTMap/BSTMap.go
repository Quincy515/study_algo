package BSTMap

import (
	"fmt"
	"reflect"
)

type Node struct {
	key         interface{}
	value       interface{}
	left, right *Node
}

type BSTMap struct {
	root *Node
	size int
}

func NewBSTMap() *BSTMap {
	return &BSTMap{}
}

func (bm *BSTMap) GetSize() int {
	return bm.size
}

func (bm *BSTMap) IsEmpty() bool {
	return bm.size == 0
}

// 向二分搜索树中添加元素(key, value)
func (bm *BSTMap) Add(key interface{}, value interface{}) {
	bm.root = bm.add(bm.root, key, value)
}

// 向以node为根的二分搜索树中插入元素(key, value),递归算法
// 返回插入新节点后二分搜索树的根
func (bm *BSTMap) add(node *Node, key interface{}, value interface{}) *Node {
	if node == nil {
		bm.size++
		return &Node{key, value, nil, nil}
	}
	if compare(key, node.key) < 0 {
		node.left = bm.add(node.left, key, value)
	} else if compare(key, node.key) > 0 {
		node.right = bm.add(node.right, key, value)
	} else { // compare(key, node.key) == 0
		node.value = value
	}
	return node
}

// 返回以node为根节点的二分搜索树中，key所在的节点
func (bm *BSTMap) getNode(node *Node, key interface{}) *Node {
	if node == nil { // 未找到等于key的节点
		return nil
	}
	if compare(key, node.key) == 0 {
		return node
	} else if compare(key, node.key) < 0 {
		return bm.getNode(node.left, key)
	} else {
		return bm.getNode(node.right, key)
	}
}

func (bm *BSTMap) Contains(key interface{}) bool {
	return bm.getNode(bm.root, key) != nil
}

func (bm *BSTMap) Get(key interface{}) interface{} {
	node := bm.getNode(bm.root, key)
	if node == nil {
		return nil
	} else {
		return node.value
	}
}

func (bm *BSTMap) Set(key interface{}, newValue interface{}) {
	node := bm.getNode(bm.root, key)
	if node == nil {
		panic(fmt.Sprintf("%v, doesn't exist.", key))
	}
	node.value = newValue
}

// 从二分搜索树中删除键为key的节点
func (bm *BSTMap) Remove(key interface{}) interface{} {
	node := bm.getNode(bm.root, key)
	if node != nil {
		bm.root = bm.remove(bm.root, key)
		return node.value
	}

	return nil
}

// 删除掉以node为根的二分搜索树中键为key的节点，递归算法
// 返回删除节点后新的二分搜索树的根
func (bm *BSTMap) remove(node *Node, key interface{}) *Node {
	if node == nil {
		return nil
	}
	if compare(key, node.key) < 0 {
		node.left = bm.remove(node.left, key)
		return node
	} else if compare(key, node.key) > 0 {
		node.right = bm.remove(node.right, key)
		return node
	} else { // compare(key, node.key)==0 删除当前节点
		// 待删除节点左子树为空的情况
		if node.left == nil {
			rightNode := node.right
			node.right = nil
			bm.size--
			return rightNode
		}
		// 待删除节点右子树为空的情况
		if node.right == nil {
			leftNode := node.left
			node.left = nil
			bm.size--
			return leftNode
		}
		// 待删除节点左右子树均不为空的情况
		// 找到比待删除节点大的最小节点，即待删除节点右子树的最小节点
		// 用这个节点顶替待删除节点的位置
		successor := bm.minimum(node.right)
		successor.right = bm.removeMin(node.right)
		successor.left = node.left

		node.left, node.right = nil, nil
		return successor
	}
}

// 返回以node为根的二分搜索树的最小值所在的节点
func (bm *BSTMap) minimum(node *Node) *Node {
	if node.left == nil {
		return node
	}
	return bm.minimum(node.left)
}

// 删除掉以node为根的二分搜索树中的最小节点
// 返回删除节点后新的二分搜索树的根
func (bm *BSTMap) removeMin(node *Node) *Node {
	if node.left == nil {
		rightNode := node.right
		node.right = nil
		bm.size--

		return rightNode
	}

	node.left = bm.removeMin(node.left)
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
