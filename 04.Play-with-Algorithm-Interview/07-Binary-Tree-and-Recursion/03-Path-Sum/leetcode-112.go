package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归方法 Time: O(n), Space: O(n)
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil { // 如果根节点为空
		return false // 返回false
	}
	// 否者如果当前根节点已经是叶子节点
	if root.Left == nil && root.Right == nil {
		return root.Val == sum // 判断节点上的值是否等于给定的和sum
	}
	// 对于其他情况则只需要递归判断左右子树即可,和sum则需要减去当前节点的值去更新
	// 左右子树只要能找到一个满足条件的路径即可
	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}

// 迭代方法 Time: O(n), Space: O(n)
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil { // 如果树为空
		return false // 返回false
	} // 定义两个栈
	var stack []*TreeNode            // 保存树节点的栈
	var sumStack []int               // 保存每个节点的和
	stack = append(stack, root)      // 树的根节点入栈
	sumStack = append(sumStack, sum) // 初始的和sum入栈

	for len(stack) > 0 { // 当栈不为空时
		n := stack[len(stack)-1] // 把两个栈的栈顶元素出栈
		stack = stack[:len(stack)-1]
		s := sumStack[len(sumStack)-1]
		sumStack = sumStack[:len(sumStack)-1]

		if n.Left == nil && n.Right == nil && n.Val == s { // 如果当前节点已经为叶子节点
			return true // 并且节点上的数字等于当前和s,则返回true
		} // 即找到一条满足条件的路径
		if n.Left != nil { // 否则如果左节点不为空
			stack = append(stack, n.Left)        // 则把左节点入栈
			sumStack = append(sumStack, s-n.Val) // 同时把更新后的和入到sumStack栈中
		}
		if n.Right != nil { // 同理对于右节点也做同样操作，如果右节点不为空
			stack = append(stack, n.Right)       // 则把右节点入栈
			sumStack = append(sumStack, s-n.Val) // 同时把更新后的和入到sumStack栈中
		}
	} // 如果栈已经为空，但是没有提前返回，说明没有找到一条满足条件的路径
	return false // 因此返回false
}
