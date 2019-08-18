package Inorder

import (
	"fmt"
	"testing"
)

func TestLeetCode(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}
	fmt.Println(inorderTraversal(root))
	fmt.Println(inorderTraversal2(root))
}
