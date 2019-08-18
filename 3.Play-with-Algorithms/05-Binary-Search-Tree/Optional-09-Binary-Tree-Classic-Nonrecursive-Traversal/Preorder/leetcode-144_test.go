package Preorder

import (
	"fmt"
	"testing"
)

func TestLeetCode(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}
	fmt.Println(preorderTraversal1(root))
	fmt.Println(preorderTraversal2(root))
	fmt.Println(preorderTraversal3(root))
}
