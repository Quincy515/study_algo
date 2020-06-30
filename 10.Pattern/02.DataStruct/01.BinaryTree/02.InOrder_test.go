package main

import "testing"

func TestInOrder(t *testing.T) {}

// 中序遍历 LeetCode 94
// TreeNode Definition for a binary tree node.
func inorder(root *TreeNode, output *[]int) {
	if root != nil {
		inorder(root.Left, output)
		*output = append(*output, root.Val)
		inorder(root.Right, output)
	}
}

func inorderTraversal(root *TreeNode) []int {
	var stack []*TreeNode // 辅助栈
	var res []int         // 保存结果
	for root != nil || len(stack) > 0 {
		for root != nil { // 树的左节点不为空时压入栈
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]  // 当前栈顶元素
		stack = stack[:len(stack)-1] // 弹出栈顶元素
		res = append(res, node.Val)  // 中序遍历加入结果
		root = node.Right
	}
	return res
}
