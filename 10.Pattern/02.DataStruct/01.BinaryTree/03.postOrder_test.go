package main

import "testing"

// 后序遍历 LeetCode 145
func TestPostOrder(t *testing.T) {}

func postorder(root *TreeNode, output *[]int) {
	if root != nil {
		postorder(root.Left, output)
		postorder(root.Right, output)
		*output = append(*output, root.Val)
	}
}

// 入栈当前节点，然后不断访问左子树
// 当左子树为空时，将它移动到栈顶节点的右节点
// 记录上一个出栈的节点变量 pre, 说明已经处理过
func postorderTraversal(root *TreeNode) []int {
	var res []int         // 定义结果列表
	var stack []*TreeNode // 定义辅助栈
	pre := new(TreeNode)  // 记录上一次出栈的节点

	for root != nil || len(stack) > 0 {
		if root != nil { // 当前节点不为空
			stack = append(stack, root) // 入栈
			root = root.Left            // 移动到它的左子树
		} else { // 当前节点为空，看栈顶节点的右节点
			root = stack[len(stack)-1].Right
			if root == nil || root == pre { // 为空，或已出栈
				pre = stack[len(stack)-1]    // 把栈顶元素存储到pre
				stack = stack[:len(stack)-1] // 弹出栈顶元素
				res = append(res, pre.Val)   // 加入到结果列表
				root = nil                   // 避免因为root等于pre而进入这条语句
			}
		}
	}
	return res
}

/**
注意点  根节点必须在右节点弹出之后，再弹出
*/
