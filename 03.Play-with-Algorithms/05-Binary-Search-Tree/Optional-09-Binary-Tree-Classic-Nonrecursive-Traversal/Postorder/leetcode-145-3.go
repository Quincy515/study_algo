package Postorder

// Non-Recursive
// Using two stacks, Reverse Preorder Traversal!
//
// Time Complexity: O(n)
// Space Complexity: O(n)
func postorderTraversal3(root *TreeNode) []int {
	var stack []*TreeNode
	var output []*TreeNode

	p := root
	for p != nil || len(stack) > 0 {
		if p != nil {
			stack = append(stack, p)
			output = append(output, p)
			p = p.Right
		} else {
			p = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			p = p.Left
		}
	}

	var res []int
	for len(output) > 0 {
		res = append(res, output[len(output)-1].Val)
		output = output[:len(output)-1]
	}
	return res
}
