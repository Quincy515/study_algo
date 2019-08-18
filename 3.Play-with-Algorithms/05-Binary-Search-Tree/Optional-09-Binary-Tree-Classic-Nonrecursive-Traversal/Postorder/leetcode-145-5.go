package Postorder

// Classic Non-Recursive
// Using a pre pointer to record the last visited node
//
// Time Complexity: O(n)
// Space Complexity: O(n)
func postorderTraversal5(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	var stack []*TreeNode
	pre := &TreeNode{}
	cur := root

	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cur.Right == nil || pre == cur.Right {
			res = append(res, cur.Val)
			pre = cur
			cur = nil
		} else {
			stack = append(stack, cur)
			cur = cur.Right
		}
	}
	return res
}
