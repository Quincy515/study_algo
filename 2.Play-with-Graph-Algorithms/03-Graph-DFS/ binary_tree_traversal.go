/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	var result []int
	preorder(root, &result)
	return result
}

func preorder(root *TreeNode, output *[]int) { // 前序遍历
	if root != nil {
		*output = append(*output, root.Val)
		preorder(root.Left, output)
		preorder(root.Right, output)
	}
}

func postorder(root *TreeNode, output *[]int) { // 后序遍历
	if root != nil {
		postorder(root.Left, output)
		postorder(root.Right, output)
		*output = append(*output, root.Val)
	}
}

func inorder(root *TreeNode, output *[]int) { // 中序遍历
	if root != nil {
		inorder(root.Left, output)
		*output = append(*output, root.Val)
		inorder(root.Right, output)
	}
}
