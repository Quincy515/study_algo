package Postorder

import (
	"fmt"
	"testing"
)

func TestLeetCode(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}
	fmt.Println(postorderTraversal1(root))
	fmt.Println(postorderTraversal2(root))
	fmt.Println(postorderTraversal3(root))
	fmt.Println(postorderTraversal4(root))
	fmt.Println(postorderTraversal5(root))
	fmt.Println(postorderTraversal6(root))
}
