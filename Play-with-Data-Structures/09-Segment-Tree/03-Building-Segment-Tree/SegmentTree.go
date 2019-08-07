package _3_Building_Segment_Tree

import (
	"bytes"
	"fmt"
)

type SegmentTree struct {
	data   []interface{}
	tree   []interface{}
	merger func(interface{}, interface{}) interface{}
}

func NewSegmentTree(arr []interface{}, merger func(interface{}, interface{}) interface{}) *SegmentTree {
	segmentTree := &SegmentTree{
		data:   make([]interface{}, len(arr)),
		tree:   make([]interface{}, len(arr)*4),
		merger: merger,
	}
	for i := 0; i < len(arr); i++ {
		segmentTree.data[i] = arr[i]
	}
	segmentTree.buildSegmentTree(0, 0, len(arr)-1)
	return segmentTree
}

// 在treeIndex的位置创建表示区间[l...r]的线段树
// treeIndex: 创建的线段树根节点索引, l,r: 该节点表示的区间
func (st *SegmentTree) buildSegmentTree(treeIndex, l, r int) {
	if l == r { // 区间长度为1，只有一个元素
		st.tree[treeIndex] = st.data[l] // 该节点存储的就是元素本身
		return
	}

	leftTreeIndex := leftChild(treeIndex)
	rightTreeIndex := rightChild(treeIndex)

	mid := l + (r-l)/2
	st.buildSegmentTree(leftTreeIndex, l, mid)
	st.buildSegmentTree(rightTreeIndex, mid+1, r)

	st.tree[treeIndex] = st.merger(st.tree[leftTreeIndex], st.tree[rightTreeIndex])
}

func (st *SegmentTree) GetSize() int {
	return len(st.data)
}

func (st *SegmentTree) Get(index int) interface{} {
	if index < 0 || index >= len(st.data) {
		panic("Index is illegal.")
	}
	return st.data[index]
}

// 返回完全二叉树的数组表示中，一个索引所表示的元素的左孩子节点的索引
func leftChild(index int) int {
	return index*2 + 1
}

// 返回完全二叉树的数组表示中，一个索引所表示的元素的右孩子节点的索引
func rightChild(index int) int {
	return index*2 + 2
}

func (st *SegmentTree) String() string {
	buffer := bytes.Buffer{}

	buffer.WriteString("[")
	for i := 0; i < len(st.tree); i++ {
		if st.tree[i] != nil {
			buffer.WriteString(fmt.Sprint(st.tree[i]))
		} else {
			buffer.WriteString("nil")
		}

		if i != len(st.tree)-1 {
			buffer.WriteString(", ")
		}
	}

	buffer.WriteString("]")
	return buffer.String()
}
