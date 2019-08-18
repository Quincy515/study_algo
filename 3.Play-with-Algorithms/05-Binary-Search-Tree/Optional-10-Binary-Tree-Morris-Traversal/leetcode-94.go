package Optional_10_Binary_Tree_Morris_Traversal

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 迭代实现
func inorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	cur := root
	for cur != nil {
		if cur.Left == nil {
			res = append(res, cur.Val)
			cur = cur.Right
		} else {
			prev := cur.Left
			for prev.Right != nil && prev.Right != cur {
				prev = prev.Right
			}
			if prev.Right == nil {
				prev.Right = cur
				cur = cur.Left
			} else {
				prev.Right = nil
				res = append(res, cur.Val)
				cur = cur.Right
			}
		}
	}
	return res
}

// 递归实现
func inorderTraversal2(root *TreeNode) []int {
	var result []int
	inorder(root, &result)
	return result
}

func inorder(root *TreeNode, output *[]int) {
	if root != nil {
		inorder(root.Left, output)
		*output = append(*output, root.Val)
		inorder(root.Right, output)
	}
}
