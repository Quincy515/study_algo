package Optional_10_Binary_Tree_Morris_Traversal

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 代码逻辑有问题，还没有解决
func postorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	dummyRoot := &TreeNode{Val: -1}
	dummyRoot.Left = root

	cur := dummyRoot
	for cur != nil {
		if cur.Left == nil {
			cur = cur.Right
		} else {
			pre := cur.Left
			for pre.Right != nil && pre.Right != cur {
				pre = pre.Right
			}

			if pre.Right == nil {
				pre.Right = cur
				cur = cur.Left
			} else {
				pre.Right = nil
				reverseTraversal(cur.Left, res)
				cur = cur.Right
			}
		}
	}
	dummyRoot = nil

	return res
}

func reverseTraversal(node *TreeNode, res []int) {
	start := len(res)
	for node != nil {
		res = append(res, node.Val)
		node = node.Right
	}

	i, j := start, len(res)-1
	for i < j {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
}

// leetcode官方题解
// https://leetcode-cn.com/problems/binary-tree-postorder-traversal/solution/er-cha-shu-de-hou-xu-bian-li-by-leetcode/
func postorderTraversal2(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	output := make([]int, 0)
	if root == nil {
		return output
	}
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		output = append(output, node.Val)
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	reverse(output)
	return output
}

func reverse(arr []int) {
	i, j := 0, len(arr)-1
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}

func main() {
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}
	fmt.Println(postorderTraversal2(root))
}
