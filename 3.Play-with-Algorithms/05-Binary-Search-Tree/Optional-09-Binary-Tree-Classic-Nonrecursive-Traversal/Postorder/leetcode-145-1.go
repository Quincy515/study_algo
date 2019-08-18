package Postorder

type TagNode struct {
	node    *TreeNode
	isFirst bool
}

// Non-Recursive
// Using a tag to record whether the node has been visited
//
// Time Complexity: O(n), n is the node number in the tree
// Space Complexity: O(h), h is the height of the tree
func postorderTraversal1(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	var stack []*TagNode
	cur := root
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, &TagNode{cur, false})
			cur = cur.Left
		}
		tagNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		cur = tagNode.node
		if tagNode.isFirst == false {
			tagNode.isFirst = true
			stack = append(stack, tagNode)
			cur = cur.Right
		} else {
			res = append(res, cur.Val)
			cur = nil
		}
	}
	return res
}
