package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// Time: O(h), Space: O(h) h是树的高度
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil { // 如果当前节点为空
		return nil // 就返回空指针
	}
	if key < root.Val { // 如果要删除的值小于根节点的值
		// 就递归去左子树上删除节点，并且返回的二叉树是当前根节点的左子树
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val { // 如果要删除的值大于根节点的值
		// 就递归去右子树上删除节点，并且返回的二叉树是当前根节点的右子树
		root.Right = deleteNode(root.Right, key)
	} else { // 如果不是以上两种情况，说明要删除的值等于当前根节点的值
		if root.Left == nil { // 先判断当前根节点是否只有一个子树
			return root.Right // 是的话就返回那个唯一子树即可
		} else if root.Right == nil {
			return root.Left
		}
		// 否则说明左右子树都不为空,去找左子树中的最大值对应的节点
		leftMax := root.Left       // 先移动到左子树
		for leftMax.Right != nil { // 如果左子树的右子树不为空
			leftMax = leftMax.Right // 就一直移动到它的右子树直到为空
		} // 找到左子树的最大值节点后，把当前根节点的右子树作为它的右子树
		leftMax.Right = root.Right
		root = root.Left // 然后把根节点移动到左子树
	}
	return root // 最后返回当前根节点即可
}
