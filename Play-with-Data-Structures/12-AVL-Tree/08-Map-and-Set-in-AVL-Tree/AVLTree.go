package AVLMap

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

// 判断该二叉树是否是一棵二分搜索树
func (avl *AVLTree) IsBST() bool {
	var keys []interface{}
	inOrder(avl.root, keys)

	for i := 1; i < len(keys); i++ {
		if compare(keys[i-1], keys[i]) > 0 {
			return false
		}
	}
	return true
}

func inOrder(node *Node, keys []interface{}) {
	if node == nil {
		return
	}
	inOrder(node.left, keys)
	keys = append(keys, node.key)
	inOrder(node.right, keys)
}

// 判断该二叉树是否是一棵平衡二叉树-左右子树高度差不超过1
func (avl *AVLTree) IsBalanced() bool {
	return avl.isBalanced(avl.root)
}

// 判断以Node为根的二叉树是否是一棵平衡二叉树，递归算法
func (avl *AVLTree) isBalanced(node *Node) bool {
	if node == nil {
		return true
	}

	balanceFactor := avl.getBalanceFactor(node)
	if math.Abs(float64(balanceFactor)) > 1 {
		return false
	}
	return avl.isBalanced(node.left) && avl.isBalanced(node.right)
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

// 对节点y进行向右旋转操作，返回旋转后新的根节点x
//        y                              x
//       / \                           /   \
//      x   T4     向右旋转 (y)        z     y
//     / \       - - - - - - - ->    / \   / \
//    z   T3                       T1  T2 T3 T4
//   / \
// T1   T2
func (avl *AVLTree) rightRotate(y *Node) *Node {
	x := y.left
	T3 := x.right

	// 向右选择过程
	x.right = y
	y.left = T3

	// 更新height
	y.height = int(math.Max(float64(avl.getHeight(y.left)), float64(avl.getHeight(y.right)))) + 1
	x.height = int(math.Max(float64(avl.getHeight(x.left)), float64(avl.getHeight(x.right)))) + 1

	return x
}

// 对节点y进行向左旋转操作，返回旋转后新的根节点x
//    y                             x
//  /  \                          /   \
// T1   x      向左旋转 (y)       y     z
//     / \   - - - - - - - ->   / \   / \
//   T2  z                     T1 T2 T3 T4
//      / \
//     T3 T4
func (avl *AVLTree) leftRotate(y *Node) *Node {
	x := y.right
	T2 := x.left

	// 向左旋转过程
	x.left = y
	y.right = T2

	// 更新height
	y.height = int(math.Max(float64(avl.getHeight(y.left)), float64(avl.getHeight(y.right)))) + 1
	x.height = int(math.Max(float64(avl.getHeight(x.left)), float64(avl.getHeight(x.right)))) + 1

	return x
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
	//if math.Abs(float64(balanceFactor)) > 1 {
	//	fmt.Println("unbalanced: ", balanceFactor)
	//}

	// 平衡维护
	// LL
	if balanceFactor > 1 && avl.getBalanceFactor(node.left) >= 0 {
		return avl.rightRotate(node)
	}
	// RR
	if balanceFactor < -1 && avl.getBalanceFactor(node.right) <= 0 {
		return avl.leftRotate(node)
	}
	// LR
	if balanceFactor > 1 && avl.getBalanceFactor(node.left) < 0 {
		node.left = avl.leftRotate(node.left)
		return avl.rightRotate(node)
	}
	// RL
	if balanceFactor < -1 && avl.getBalanceFactor(node.right) > 0 {
		node.right = avl.rightRotate(node.right)
		return avl.leftRotate(node)
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

	var retNode *Node
	if compare(key, node.key) < 0 {
		node.left = avl.remove(node.left, key)
		retNode = node
	} else if compare(key, node.key) > 0 {
		node.right = avl.remove(node.right, key)
		retNode = node
	} else { // compare(key, node.key)==0 删除当前节点
		// 待删除节点左子树为空的情况
		if node.left == nil {
			rightNode := node.right
			node.right = nil
			avl.size--
			retNode = rightNode
		} else if node.right == nil {
			// 待删除节点右子树为空的情况
			leftNode := node.left
			node.left = nil
			avl.size--
			retNode = leftNode
		} else {
			// 待删除节点左右子树均不为空的情况
			// 找到比待删除节点大的最小节点，即待删除节点右子树的最小节点
			// 用这个节点顶替待删除节点的位置
			successor := avl.minimum(node.right)
			successor.right = avl.remove(node.right, successor.key)
			successor.left = node.left

			node.left, node.right = nil, nil
			retNode = successor
		}
	}

	// 删除一个节点之后，平衡的维护
	if retNode == nil {
		return nil
	}
	// 更新height
	retNode.height = 1 + int(math.Max(
		float64(avl.getHeight(retNode.left)),
		float64(avl.getHeight(retNode.right))))
	// 计算平衡因子
	balanceFactor := avl.getBalanceFactor(retNode)

	// 平衡维护
	// LL
	if balanceFactor > 1 && avl.getBalanceFactor(retNode.left) >= 0 {
		return avl.rightRotate(retNode)
	}
	// RR
	if balanceFactor < -1 && avl.getBalanceFactor(retNode.right) <= 0 {
		return avl.leftRotate(retNode)
	}
	// LR
	if balanceFactor > 1 && avl.getBalanceFactor(retNode.left) < 0 {
		retNode.left = avl.leftRotate(retNode.left)
		return avl.rightRotate(retNode)
	}
	// RL
	if balanceFactor < -1 && avl.getBalanceFactor(retNode.right) > 0 {
		retNode.right = avl.rightRotate(retNode.right)
		return avl.leftRotate(retNode)
	}
	return retNode
}

// 返回以node为根的二分搜索树的最小值所在的节点
func (avl *AVLTree) minimum(node *Node) *Node {
	if node.left == nil {
		return node
	}
	return avl.minimum(node.left)
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
