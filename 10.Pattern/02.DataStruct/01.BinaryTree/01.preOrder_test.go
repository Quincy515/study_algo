package main

import "testing"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
二叉树遍历
- 前序遍历：先访问根节点，再前序遍历左子树，再前序遍历右子树
- 中序遍历：先中序遍历左子树，再访问根节点，再中序遍历右子树
- 后序遍历：先后序遍历左子树，再后序遍历右子树，再访问根节点

注意点：
- 以根访问顺序决定是什么遍历
- 左子树都是优先右子树
*/

// 前序递归 LeetCode 144
func preorder(root *TreeNode, output *[]int) {
	if root != nil {
		// 先访文根节点再访问左右子树
		*output = append(*output, root.Val)
		preorder(root.Left, output)
		preorder(root.Right, output)
	}
}

func TestBinaryTreeTraverse(t *testing.T) {
	var (
		root *TreeNode
		res  []int
	)
	preorder(root, &res)

}

// 前序遍历-迭代方法  Time: O(n) Space: O(n)
func preorderTraversal(root *TreeNode) []int {
	// 终止条件
	if root == nil {
		return nil
	}
	var res []int
	var stack []*TreeNode // 用来暂存节点
	for root != nil || len(stack) != 0 {
		for root != nil { // 节点不为空，把节点的值加入到结果res
			// 前序遍历，所以先保存结果
			res = append(res, root.Val)
			stack = append(stack, root) //节点压入栈
			root = root.Left            // 移动到它的左节点
		}
		// pop
		node := stack[len(stack)-1] //节点为空时，出栈
		stack = stack[:len(stack)-1]
		root = node.Right // 并把它移动到右节点
	}
	return res
}
