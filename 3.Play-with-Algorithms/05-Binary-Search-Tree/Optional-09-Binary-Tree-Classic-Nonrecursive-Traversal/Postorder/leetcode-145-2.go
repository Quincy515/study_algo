package Postorder

// Non-Recursive
// Using two stacks, Reverse Preorder Traversal!
//
// Time Complexity: O(n)
// Space Complexity: O(n)
func postorderTraversal2(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	var stack []*TreeNode
	var output []int

	stack = append(stack, root)
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		output = append(output, cur.Val)

		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
	}

	for len(output) > 0 {
		res = append(res, output[len(output)-1])
		output = output[:len(output)-1]
	}
	return res
}
