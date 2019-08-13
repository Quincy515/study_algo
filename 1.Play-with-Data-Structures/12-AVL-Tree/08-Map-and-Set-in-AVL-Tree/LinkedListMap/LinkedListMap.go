package LinkedListMap

import (
	"fmt"
)

type Node struct {
	key, value interface{}
	next       *Node
}

func (n *Node) String() string {
	return fmt.Sprint(n.key, " : ", n.value)
}

type LinkedListMap struct {
	dummyHead *Node
	size      int
}

func NewLinkedListMap() *LinkedListMap {
	return &LinkedListMap{&Node{}, 0}
}

func (lm *LinkedListMap) GetSize() int {
	return lm.size
}

func (lm *LinkedListMap) IsEmpty() bool {
	return lm.size == 0
}

// 传入一个key返回该节点的引用
func (lm *LinkedListMap) getNode(key interface{}) *Node {
	cur := lm.dummyHead.next
	for cur != nil {
		if cur.key == key {
			return cur
		}
		cur = cur.next
	}
	return nil
}

func (lm *LinkedListMap) Contains(key interface{}) bool {
	return lm.getNode(key) != nil
}

func (lm *LinkedListMap) Get(key interface{}) interface{} {
	node := lm.getNode(key)
	if node == nil {
		return nil
	} else {
		return node.value
	}
}

func (lm *LinkedListMap) Add(key interface{}, value interface{}) {
	node := lm.getNode(key)
	if node == nil {
		lm.dummyHead.next = &Node{key, value, lm.dummyHead.next}
		lm.size++
	} else { // 需要添加的键值已经存在，则更新，也可以抛出异常
		node.value = value
	}
}

func (lm *LinkedListMap) Set(key interface{}, newValue interface{}) {
	node := lm.getNode(key)
	if node == nil {
		panic(fmt.Sprintf("%v, doesn't exist", key))
	}
	node.value = newValue
}

func (lm *LinkedListMap) Remove(key interface{}) interface{} {
	// RemoveElement代码逻辑修改
	prev := lm.dummyHead
	for prev.next != nil {
		if prev.next.key == key {
			break // 寻找到要删除的节点，跳出
		}
		prev = prev.next
	}

	if prev.next != nil {
		delNode := prev.next
		prev.next = delNode.next
		delNode.next = nil
		lm.size--
		return delNode.value
	}
	return nil
}
