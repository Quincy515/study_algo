package Postorder

// Non-Recursive
// Using a pre pointer to record the last visited node
//
// Time Complexity: O(n)
// Space Complexity: O(n)
func postorderTraversal4(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	var stack []*TreeNode
	pre := &TreeNode{}

	stack = append(stack, root)
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if (cur.Left == nil && cur.Right == nil) ||
			(pre != nil && pre == cur.Left && cur.Right == nil) ||
			(pre != nil && pre == cur.Right) {
			res = append(res, cur.Val)
			pre = cur
		} else {
			stack = append(stack, cur)
			if cur.Right != nil {
				stack = append(stack, cur.Right)
			}
			if cur.Left != nil {
				stack = append(stack, cur.Left)
			}
		}
	}

	return res
}
