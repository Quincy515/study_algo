package Optional_10_Binary_Tree_Morris_Traversal

func preorderTraversal(root *TreeNode) []int {
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
				res = append(res, cur.Val)
				prev.Right = cur
				cur = cur.Left
			} else {
				prev.Right = nil
				cur = cur.Right
			}
		}
	}
	return res
}
